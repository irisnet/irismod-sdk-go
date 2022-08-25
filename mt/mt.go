package mt

import (
	"context"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
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

// MintMT create new MT
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
		DenomId:   request.DenomID,
		Amount:    request.Amount,
		Data:      request.Data,
		Sender:    sender.String(),
		Recipient: recipient,
	}
	return mc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

// AddMT issuing additional MT
func (mc mtClient) AddMT(request AddMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
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
		DenomId:   request.DenomID,
		Amount:    request.Amount,
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
		DenomId: request.DenomID,
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
		DenomId:   request.DenomID,
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
		DenomId: request.DenomID,
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

func (mc mtClient) QuerySupply(denomID, creator string) (uint64, sdk.Error) {
	if len(denomID) == 0 {
		return 0, sdk.Wrapf("denomID is required")
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
			DenomId: denomID,
		},
	)
	if err != nil {
		return 0, sdk.Wrap(err)
	}

	return res.Amount, nil
}

func (mc mtClient) QueryMTSupply(denomID, mtID string) (uint64, sdk.Error) {
	if len(denomID) == 0 {
		return 0, sdk.Wrapf("denomID is required")
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
			DenomId: denomID,
			MtId:    mtID,
		},
	)
	if err != nil {
		return 0, sdk.Wrap(err)
	}

	return res.Amount, nil
}

func (mc mtClient) QueryDenom(denomID string) (QueryDenomResp, sdk.Error) {
	if len(denomID) == 0 {
		return QueryDenomResp{}, sdk.Wrapf("denomID is required")
	}

	conn, err := mc.GenConn()
	if err != nil {
		return QueryDenomResp{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Denom(
		context.Background(),
		&QueryDenomRequest{DenomId: denomID},
	)
	if err != nil {
		return QueryDenomResp{}, sdk.Wrap(err)
	}

	return res.Denom.Convert().(QueryDenomResp), nil
}

func (mc mtClient) QueryDenoms(pageReq *query.PageRequest) ([]QueryDenomResp, sdk.Error) {
	conn, err := mc.GenConn()

	if err != nil {
		return nil, sdk.Wrap(err)
	}

	pagination, e := query.FormatPageRequest(pageReq)
	if e != nil {
		return nil, e
	}

	res, err := NewQueryClient(conn).Denoms(
		context.Background(),
		&QueryDenomsRequest{
			Pagination: pagination,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	return denoms(res.Denoms).Convert().([]QueryDenomResp), nil
}

func (mc mtClient) QueryMT(denomID, mtID string) (QueryMTResp, sdk.Error) {
	if len(denomID) == 0 {
		return QueryMTResp{}, sdk.Wrapf("denomID is required")
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
			DenomId: denomID,
			MtId:    mtID,
		},
	)
	if err != nil {
		return QueryMTResp{}, sdk.Wrap(err)
	}

	return res.Mt.Convert().(QueryMTResp), nil
}

func (mc mtClient) QueryMTs(denomID string, pageReq *query.PageRequest) ([]QueryMTResp, sdk.Error) {
	if len(denomID) == 0 {
		return nil, sdk.Wrapf("denomID is required")
	}

	conn, err := mc.GenConn()

	if err != nil {
		return nil, sdk.Wrap(err)
	}

	pagination, e := query.FormatPageRequest(pageReq)
	if e != nil {
		return nil, e
	}

	res, err := NewQueryClient(conn).MTs(
		context.Background(),
		&QueryMTsRequest{
			DenomId:    denomID,
			Pagination: pagination,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}

	return mts(res.Mts).Convert().([]QueryMTResp), nil
}

func (mc mtClient) QueryBalances(denomID, owner string) ([]QueryBalanceResp, sdk.Error) {
	if len(denomID) == 0 {
		return nil, sdk.Wrapf("denomID is required")
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
			DenomId: denomID,
			Owner:   owner,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	return balances(res.Balance).Convert().([]QueryBalanceResp), nil
}
