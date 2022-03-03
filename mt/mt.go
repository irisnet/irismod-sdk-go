package mt

import (
	"context"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type mtClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return mtClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (mtClient) Name() string {
	return ModuleName
}

func (mtClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (mc mtClient) IssueDenom(request IssueDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := mc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgIssueDenom{
		Name:   request.Name,
		Sender: sender.String(),
		Data:   request.Data,
	}
	return mc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (mc mtClient) MintMT(request MintMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := mc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	var recipient = sender.String()
	if len(request.Recipient) > 0 {
		if err := sdk.ValidateAccAddress(request.Recipient); err != nil {
			return sdk.ResultTx{}, sdk.Wrap(err)
		}
		recipient = request.Recipient
	}

	msg := &MsgMintMT{
		Id:        request.ID,
		DenomId:   request.Denom,
		Amount:    request.Amount,
		Data:      request.Data,
		Sender:    sender.String(),
		Recipient: recipient,
	}
	return mc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (mc mtClient) EditMT(request EditMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := mc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgEditMT{
		Id:      request.ID,
		DenomId: request.Denom,
		Data:    request.Data,
		Sender:  sender.String(),
	}
	return mc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (mc mtClient) TransferMT(request TransferMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := mc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	if err := sdk.ValidateAccAddress(request.Recipient); err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgTransferMT{
		Id:        request.ID,
		DenomId:   request.Denom,
		Amount:    request.Amount,
		Sender:    sender.String(),
		Recipient: request.Recipient,
	}
	return mc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (mc mtClient) BurnMT(request BurnMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := mc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgBurnMT{
		Id:      request.ID,
		DenomId: request.Denom,
		Amount:  request.Amount,
		Sender:  sender.String(),
	}
	return mc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (mc mtClient) TransferDenom(request TransferDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := mc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	if err := sdk.ValidateAccAddress(request.Recipient); err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgTransferDenom{
		Id:        request.ID,
		Sender:    sender.String(),
		Recipient: request.Recipient,
	}
	return mc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (mc mtClient) QuerySupply(denom, creator string) (uint64, sdk.Error) {
	if len(denom) == 0 {
		return 0, sdk.Wrapf("denom is required")
	}

	if err := sdk.ValidateAccAddress(creator); err != nil {
		return 0, sdk.Wrap(err)
	}

	conn, err := mc.GenConn()

	if err != nil {
		return 0, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Supply(
		context.Background(),
		&QuerySupplyRequest{
			Owner:   creator,
			DenomId: denom,
		},
	)
	if err != nil {
		return 0, sdk.Wrap(err)
	}

	return res.Amount, nil
}

func (mc mtClient) QueryMTSupply(denom, mtID string) (uint64, sdk.Error) {
	if len(denom) == 0 {
		return 0, sdk.Wrapf("denom is required")
	}
	if len(mtID) == 0 {
		return 0, sdk.Wrapf("mtID is required")
	}

	conn, err := mc.GenConn()

	if err != nil {
		return 0, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).MTSupply(
		context.Background(),
		&QueryMTSupplyRequest{
			DenomId: denom,
			MtId:    mtID,
		},
	)
	if err != nil {
		return 0, sdk.Wrap(err)
	}

	return res.Amount, nil
}

func (mc mtClient) QueryDenom(denom string) (QueryDenomResp, sdk.Error) {
	if len(denom) == 0 {
		return QueryDenomResp{}, sdk.Wrapf("denom is required")
	}

	conn, err := mc.GenConn()
	if err != nil {
		return QueryDenomResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Denom(
		context.Background(),
		&QueryDenomRequest{DenomId: denom},
	)
	if err != nil {
		return QueryDenomResp{}, sdk.Wrap(err)
	}

	return res.Denom.Convert().(QueryDenomResp), nil
}

func (mc mtClient) QueryDenoms() ([]QueryDenomResp, sdk.Error) {
	conn, err := mc.GenConn()

	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Denoms(
		context.Background(),
		&QueryDenomsRequest{},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	return denoms(res.Denoms).Convert().([]QueryDenomResp), nil
}

func (mc mtClient) QueryMT(denom, mtID string) (QueryMTResp, sdk.Error) {
	if len(denom) == 0 {
		return QueryMTResp{}, sdk.Wrapf("denom is required")
	}

	if len(mtID) == 0 {
		return QueryMTResp{}, sdk.Wrapf("mtID is required")
	}

	conn, err := mc.GenConn()

	if err != nil {
		return QueryMTResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).MT(
		context.Background(),
		&QueryMTRequest{
			DenomId: denom,
			MtId:    mtID,
		},
	)
	if err != nil {
		return QueryMTResp{}, sdk.Wrap(err)
	}

	return res.Mt.Convert().(QueryMTResp), nil
}

func (mc mtClient) QueryMTs(denom string) ([]QueryMTResp, sdk.Error) {
	if len(denom) == 0 {
		return nil, sdk.Wrapf("denom is required")
	}

	conn, err := mc.GenConn()

	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).MTs(
		context.Background(),
		&QueryMTsRequest{
			DenomId: denom,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	return mts(res.Mts).Convert().([]QueryMTResp), nil
}

func (mc mtClient) QueryBalances(denom, owner string) ([]QueryBalanceResp, sdk.Error) {
	if len(denom) == 0 {
		return nil, sdk.Wrapf("denom is required")
	}

	if err := sdk.ValidateAccAddress(owner); err != nil {
		return nil, sdk.Wrap(err)
	}

	conn, err := mc.GenConn()
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Balances(
		context.Background(),
		&QueryBalancesRequest{
			DenomId: denom,
			Owner:   owner,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return balances(res.Balance).Convert().([]QueryBalanceResp), nil
}
