package coinswap

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/irisnet/core-sdk-go/types/query"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	ctypes "github.com/irisnet/core-sdk-go/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
)

type coinswapClient struct {
	sdk.BaseClient
	codec.Codec
	totalSupply
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec, queryTotalSupply totalSupply) Client {
	return coinswapClient{
		BaseClient:  bc,
		Codec:       cdc,
		totalSupply: queryTotalSupply,
	}
}

func (swap coinswapClient) Name() string {
	return ModuleName
}

func (swap coinswapClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (swap coinswapClient) AddLiquidity(request AddLiquidityRequest, baseTx sdk.BaseTx) (*AddLiquidityResponse, error) {
	creator, err := swap.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrQueryAddress, err.Error())
	}

	msg := &MsgAddLiquidity{
		MaxToken:         request.MaxToken,
		ExactStandardAmt: ctypes.NewInt(request.BaseAmt.Int64()),
		MinLiquidity:     ctypes.NewInt(request.MinLiquidity.Int64()),
		Deadline:         request.Deadline,
		Sender:           creator.String(),
	}

	res, err := swap.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return nil, err
	}
	var totalCoins = sdk.NewCoins()
	coinStrs := sdk.StringifyEvents(res.TxResult.Events).GetValues(eventTypeTransfer, attributeKeyAmount)
	if len(coinStrs) != 3 {
		return nil, sdkerrors.Wrap(ErrLiquidityPool, fmt.Sprintf("coinstrs lenght is %d want 3", len(coinStrs)))
	}
	if !strings.Contains(coinStrs[2], "lpt-") {
		return nil, sdkerrors.Wrap(ErrLiquidityPool, "lpt not found ")
	}
	liquidityDenom := "lpt-" + strings.Split(coinStrs[2], "-")[1]
	for _, coinStr := range coinStrs {
		coins, er := sdk.ParseCoinsNormalized(coinStr)
		if er != nil {
			swap.Logger().Error("Parse coin str failed", "coin", coinStr)
			continue
		}
		totalCoins = totalCoins.Add(coins...)
	}
	response := &AddLiquidityResponse{
		TokenAmt:  totalCoins.AmountOf(request.MaxToken.Denom),
		BaseAmt:   request.BaseAmt,
		Liquidity: totalCoins.AmountOf(liquidityDenom),
		TxHash:    res.Hash.String(),
	}
	return response, nil
}

func (swap coinswapClient) RemoveLiquidity(request RemoveLiquidityRequest,
	baseTx sdk.BaseTx) (*RemoveLiquidityResponse, error) {
	creator, err := swap.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrQueryAddress, err.Error())
	}

	msg := &MsgRemoveLiquidity{
		WithdrawLiquidity: request.Liquidity,
		MinToken:          ctypes.NewInt(request.MinTokenAmt.Int64()),
		MinStandardAmt:    ctypes.NewInt(request.MinBaseAmt.Int64()),
		Deadline:          request.Deadline,
		Sender:            creator.String(),
	}

	res, err := swap.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return nil, err
	}

	var totalCoins = sdk.NewCoins()
	coinStrs := sdk.StringifyEvents(res.TxResult.Events).GetValues(eventTypeTransfer, attributeKeyAmount)
	for _, coinStr := range coinStrs {
		coins, er := sdk.ParseCoinsNormalized(coinStr)
		if er != nil {
			swap.Logger().Error("Parse coin str failed", "coin", coinStr)
			continue
		}
		totalCoins = totalCoins.Add(coins...)
	}

	tokenDenom, er := GetTokenDenomFrom(request.Liquidity.Denom)
	if er != nil {
		return nil, er
	}

	response := &RemoveLiquidityResponse{
		TokenAmt:  totalCoins.AmountOf(tokenDenom),
		BaseAmt:   totalCoins.AmountOf(BaseDenom),
		Liquidity: request.Liquidity,
		TxHash:    res.Hash.String(),
	}
	return response, nil
}

func (swap coinswapClient) SwapCoin(request SwapCoinRequest, baseTx sdk.BaseTx) (*SwapCoinResponse, error) {
	creator, err := swap.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	input := Input{
		Address: creator.String(),
		Coin:    request.Input,
	}

	if len(request.Receiver) == 0 {
		request.Receiver = input.Address
	}

	output := Output{
		Address: request.Receiver,
		Coin:    request.Output,
	}

	msg := &MsgSwapOrder{
		Input:      input,
		Output:     output,
		Deadline:   request.Deadline,
		IsBuyOrder: request.IsBuyOrder,
	}

	res, err := swap.BuildAndSend([]sdk.Msg{msg}, baseTx)
	if err != nil {
		return nil, err
	}
	amount, er := sdk.StringifyEvents(res.TxResult.Events).GetValue(eventTypeSwap, attributeKeyAmount)
	if er != nil {
		return nil, er
	}

	amt, ok := sdk.NewIntFromString(amount)
	if !ok {
		return nil, sdkerrors.Wrapf(ErrConvertInt, fmt.Sprintf("%s can not convert to sdk.Int type", amount))
	}

	inputAmt := request.Input.Amount
	outputAmt := request.Output.Amount
	if request.IsBuyOrder {
		inputAmt = amt
	} else {
		outputAmt = amt
	}

	response := &SwapCoinResponse{
		InputAmt:  inputAmt,
		OutputAmt: outputAmt,
		TxHash:    res.Hash.String(),
	}
	return response, nil
}

func (swap coinswapClient) BuyTokenWithAutoEstimate(paidTokenDenom string, boughtCoin sdk.Coin,
	deadline int64,
	baseTx sdk.BaseTx,
) (res *SwapCoinResponse, err error) {
	var amount = sdk.ZeroInt()
	switch {
	case paidTokenDenom == BaseDenom:
		amount, err = swap.EstimateBaseForBoughtToken(boughtCoin)
		break
	case boughtCoin.Denom == BaseDenom:
		amount, err = swap.EstimateTokenForBoughtBase(paidTokenDenom, boughtCoin.Amount)
		break
	default:
		amount, err = swap.EstimateTokenForBoughtToken(paidTokenDenom, boughtCoin)
		break
	}

	if err != nil {
		return nil, err
	}

	req := SwapCoinRequest{
		Input:      sdk.NewCoin(paidTokenDenom, amount),
		Output:     boughtCoin,
		Deadline:   deadline,
		IsBuyOrder: true,
	}
	return swap.SwapCoin(req, baseTx)
}

func (swap coinswapClient) SellTokenWithAutoEstimate(gotTokenDenom string, soldCoin sdk.Coin,
	deadline int64,
	baseTx sdk.BaseTx,
) (res *SwapCoinResponse, err error) {
	var amount = sdk.ZeroInt()
	switch {
	case gotTokenDenom == BaseDenom:
		amount, err = swap.EstimateBaseForSoldToken(soldCoin)
		break
	case soldCoin.Denom == BaseDenom:
		amount, err = swap.EstimateTokenForSoldBase(gotTokenDenom, soldCoin.Amount)
		break
	default:
		amount, err = swap.EstimateTokenForSoldToken(gotTokenDenom, soldCoin)
		break
	}

	if err != nil {
		return nil, err
	}

	req := SwapCoinRequest{
		Input:      soldCoin,
		Output:     sdk.NewCoin(gotTokenDenom, amount),
		Deadline:   deadline,
		IsBuyOrder: false,
	}
	return swap.SwapCoin(req, baseTx)
}

func (swap coinswapClient) QueryPool(lptDenom string) (*QueryPoolResponse, error) {
	conn, err := swap.GenConn()
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	resp, err := NewQueryClient(conn).LiquidityPool(
		context.Background(),
		&QueryLiquidityPoolRequest{LptDenom: lptDenom},
	)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrLiquidityPool, err.Error())
	}
	return resp.Convert().(*QueryPoolResponse), err
}

func (swap coinswapClient) QueryAllPools(req PageRequest) (*QueryAllPoolsResponse, error) {
	conn, err := swap.GenConn()
	resp, err := NewQueryClient(conn).LiquidityPools(
		context.Background(),
		&QueryLiquidityPoolsRequest{
			Pagination: &query.PageRequest{
				Key:        req.Key,
				Offset:     req.Offset,
				Limit:      req.Limit,
				CountTotal: req.CountTotal,
			},
		},
	)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrLiquidityPool, err.Error())
	}
	return resp.Convert().(*QueryAllPoolsResponse), err
}

func (swap coinswapClient) EstimateTokenForSoldBase(tokenDenom string,
	soldBaseAmt sdk.Int,
) (sdk.Int, error) {
	lptDenom, err := swap.GetLiquidityDenomFrom(tokenDenom)
	if err != nil {
		//TODO error
		return sdk.ZeroInt(), err
	}
	result, err := swap.QueryPool(lptDenom)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	fee := sdk.MustNewDecFromStr(result.Pool.Fee)
	amount := getInputPrice(soldBaseAmt,
		result.Pool.Standard.Amount, result.Pool.Token.Amount, fee)
	return amount, nil
}

func (swap coinswapClient) EstimateBaseForSoldToken(soldToken sdk.Coin) (sdk.Int, error) {
	lptDenom, err := swap.GetLiquidityDenomFrom(soldToken.Denom)
	if err != nil {
		//TODO error
		return sdk.ZeroInt(), err
	}
	result, err := swap.QueryPool(lptDenom)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	fee := sdk.MustNewDecFromStr(result.Pool.Fee)
	amount := getInputPrice(soldToken.Amount,
		result.Pool.Token.Amount, result.Pool.Standard.Amount, fee)
	return amount, nil
}

func (swap coinswapClient) EstimateTokenForSoldToken(boughtTokenDenom string,
	soldToken sdk.Coin) (sdk.Int, error) {
	if boughtTokenDenom == soldToken.Denom {
		return sdk.ZeroInt(), errors.New("invalid trade")
	}

	boughtBaseAmt, err := swap.EstimateBaseForSoldToken(soldToken)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	return swap.EstimateTokenForSoldBase(boughtTokenDenom, boughtBaseAmt)
}

func (swap coinswapClient) EstimateTokenForBoughtBase(soldTokenDenom string,
	exactBoughtBaseAmt sdk.Int) (sdk.Int, error) {
	lptDenom, err := swap.GetLiquidityDenomFrom(soldTokenDenom)
	if err != nil {
		//TODO error
		return sdk.ZeroInt(), err
	}
	result, err := swap.QueryPool(lptDenom)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	fee := sdk.MustNewDecFromStr(result.Pool.Fee)
	amount := getOutputPrice(exactBoughtBaseAmt,
		result.Pool.Token.Amount, result.Pool.Standard.Amount, fee)
	return amount, nil
}

func (swap coinswapClient) EstimateBaseForBoughtToken(boughtToken sdk.Coin) (sdk.Int, error) {
	lptDenom, err := swap.GetLiquidityDenomFrom(boughtToken.Denom)
	if err != nil {
		//TODO error
		return sdk.ZeroInt(), err
	}
	result, err := swap.QueryPool(lptDenom)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	fee := sdk.MustNewDecFromStr(result.Pool.Fee)
	amount := getOutputPrice(boughtToken.Amount,
		result.Pool.Standard.Amount, result.Pool.Token.Amount, fee)
	return amount, nil
}

func (swap coinswapClient) EstimateTokenForBoughtToken(soldTokenDenom string,
	boughtToken sdk.Coin) (sdk.Int, error) {
	if soldTokenDenom == boughtToken.Denom {
		return sdk.ZeroInt(), errors.New("invalid trade")
	}

	soldBaseAmt, err := swap.EstimateBaseForBoughtToken(boughtToken)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	return swap.EstimateTokenForBoughtBase(soldTokenDenom, soldBaseAmt)
}

func (swap coinswapClient) GetLiquidityDenomFrom(denom string) (string, error) {
	poolID := GetPoolId(denom)
	if denom == BaseDenom {
		return "", sdkerrors.Wrapf(ErrDenom, "should not be base denom : %s", denom)
	}
	key := fmt.Sprintf("%s/%s", "pool", poolID)
	res, err := swap.QueryStore([]byte(key), ModuleName, 0, false)
	if err != nil {
		return "", sdkerrors.Wrapf(ErrDenom, "query story : %s", key)
	}
	pool := Pool{}
	swap.Codec.MustUnmarshal(res.Value, &pool)
	return pool.LptDenom, nil
}

// GetPoolId returns the pool coin denom by specified sequence.
func GetPoolId(counterpartyDenom string) string {
	return fmt.Sprintf("pool-%s", counterpartyDenom)
}

func GetTokenDenomFrom(liquidityDenom string) (string, error) {
	if !strings.HasPrefix(liquidityDenom, "lpt") {
		return "", sdkerrors.Wrapf(ErrPrefix, "wrong liquidity denom : %s", liquidityDenom)
	}
	return strings.TrimPrefix(liquidityDenom, "lpt"), nil
}

// getInputPrice returns the amount of coins bought (calculated) given the input amount being sold (exact)
// The fee is included in the input coins being bought
// https://github.com/runtimeverification/verified-smart-contracts/blob/uniswap/uniswap/x-y-k.pdf
func getInputPrice(inputAmt, inputReserve, outputReserve sdk.Int, fee sdk.Dec) sdk.Int {
	deltaFee := sdk.OneDec().Sub(fee)
	inputAmtWithFee := inputAmt.Mul(sdk.NewIntFromBigInt(deltaFee.BigInt()))
	numerator := inputAmtWithFee.Mul(outputReserve)
	denominator := inputReserve.Mul(sdk.NewIntWithDecimal(1, sdk.Precision)).Add(inputAmtWithFee)
	return numerator.Quo(denominator)
}

// getOutputPrice returns the amount of coins sold (calculated) given the output amount being bought (exact)
// The fee is included in the output coins being bought
func getOutputPrice(outputAmt, inputReserve, outputReserve sdk.Int, fee sdk.Dec) sdk.Int {
	deltaFee := sdk.OneDec().Sub(fee)
	numerator := inputReserve.Mul(outputAmt).Mul(sdk.NewIntWithDecimal(1, sdk.Precision))
	denominator := (outputReserve.Sub(outputAmt)).Mul(sdk.NewIntFromBigInt(deltaFee.BigInt()))
	return numerator.Quo(denominator).Add(sdk.OneInt())
}
