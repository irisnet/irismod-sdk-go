package coinswap

import (
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

// expose Record module api for user
type Client interface {
	sdk.Module
	AddLiquidity(request AddLiquidityRequest,
		baseTx sdk.BaseTx) (*AddLiquidityResponse, error)
	RemoveLiquidity(request RemoveLiquidityRequest,
		baseTx sdk.BaseTx) (*RemoveLiquidityResponse, error)
	SwapCoin(request SwapCoinRequest,
		baseTx sdk.BaseTx) (*SwapCoinResponse, error)

	BuyTokenWithAutoEstimate(paidTokenDenom string, boughtCoin types.Coin,
		deadline int64,
		baseTx sdk.BaseTx,
	) (res *SwapCoinResponse, err error)
	SellTokenWithAutoEstimate(gotTokenDenom string, soldCoin types.Coin,
		deadline int64,
		baseTx sdk.BaseTx,
	) (res *SwapCoinResponse, err error)

	QueryPool(denom string) (*QueryPoolResponse, error)
	QueryAllPools() (*QueryAllPoolsResponse, error)

	EstimateTokenForSoldBase(tokenDenom string,
		soldBase math.Int,
	) (math.Int, error)
	EstimateBaseForSoldToken(soldToken types.Coin) (math.Int, error)
	EstimateTokenForSoldToken(boughtTokenDenom string,
		soldToken types.Coin) (math.Int, error)
	EstimateTokenForBoughtBase(soldTokenDenom string,
		boughtBase math.Int) (math.Int, error)
	EstimateBaseForBoughtToken(boughtToken types.Coin) (math.Int, error)
	EstimateTokenForBoughtToken(soldTokenDenom string,
		boughtToken types.Coin) (math.Int, error)
}

type AddLiquidityRequest struct {
	MaxToken     types.Coin
	BaseAmt      math.Int
	MinLiquidity math.Int
	Deadline     int64
}

type AddLiquidityResponse struct {
	TokenAmt  math.Int
	BaseAmt   math.Int
	Liquidity math.Int
	TxHash    string
}

type RemoveLiquidityRequest struct {
	MinTokenAmt math.Int
	MinBaseAmt  math.Int
	Liquidity   types.Coin
	Deadline    int64
}

type RemoveLiquidityResponse struct {
	TokenAmt  math.Int
	BaseAmt   math.Int
	Liquidity types.Coin
	TxHash    string
}

type SwapCoinRequest struct {
	Input      types.Coin
	Output     types.Coin
	Receiver   string
	Deadline   int64
	IsBuyOrder bool
}

type SwapCoinResponse struct {
	InputAmt  math.Int
	OutputAmt math.Int
	TxHash    string
}

type QueryPoolResponse struct {
	BaseCoin  types.Coin
	TokenCoin types.Coin
	Liquidity types.Coin
	Fee       string
}

type QueryAllPoolsResponse struct {
	Pools []QueryPoolResponse
}
