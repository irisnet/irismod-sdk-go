package htlc

import (
	"context"

	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type htlcClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(baseClient sdk.BaseClient, marshaler codec.Codec) Client {
	return htlcClient{
		BaseClient: baseClient,
		Codec:      marshaler,
	}
}

func (hc htlcClient) Name() string {
	return ModuleName
}

func (hc htlcClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (hc htlcClient) CreateHTLC(request CreateHTLCRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := hc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	if request.TimeLock == 0 {
		request.TimeLock = MinTimeLock
	}

	amount, err := hc.ToMinCoin(request.Amount...)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrToMintCoin, err.Error())
	}

	msg := &MsgCreateHTLC{
		Sender:               sender.String(),
		To:                   request.To,
		ReceiverOnOtherChain: request.ReceiverOnOtherChain,
		SenderOnOtherChain:   request.SenderOnOtherChain,
		Amount:               amount,
		HashLock:             request.HashLock,
		Timestamp:            request.Timestamp,
		TimeLock:             request.TimeLock,
		Transfer:             request.Transfer,
	}
	return hc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (hc htlcClient) ClaimHTLC(hashLockId string, secret string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := hc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	msg := &MsgClaimHTLC{
		Sender: sender.String(),
		Id:     hashLockId,
		Secret: secret,
	}
	return hc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (hc htlcClient) QueryHTLC(hashLockId string) (QueryHTLCResp, error) {
	if len(hashLockId) == 0 {
		return QueryHTLCResp{}, sdkerrors.Wrapf(ErrInvalidRequest, "hashLock id is required")
	}

	conn, err := hc.GenConn()
	if err != nil {
		return QueryHTLCResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).HTLC(
		context.Background(),
		&QueryHTLCRequest{
			Id: hashLockId,
		})
	if err != nil {
		return QueryHTLCResp{}, sdkerrors.Wrapf(ErrQueryHTLC, err.Error())
	}
	return res.Htlc.Convert().(QueryHTLCResp), nil
}

func (hc htlcClient) QueryParams() (QueryParamsResp, error) {

	conn, err := hc.GenConn()
	if err != nil {
		return QueryParamsResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Params(
		context.Background(),
		&QueryParamsRequest{})
	if err != nil {
		return QueryParamsResp{}, sdkerrors.Wrapf(ErrQueryParams, err.Error())
	}
	return res.Params.Convert().(QueryParamsResp), nil
}
