package random

import (
	"context"
	"strconv"

	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	cdctypes "github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type randomClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(baseClient sdk.BaseClient, marshaler codec.Codec) *randomClient {
	return &randomClient{
		BaseClient: baseClient,
		Codec:      marshaler,
	}
}

func (rc randomClient) Name() string {
	return ModuleName
}

func (rc randomClient) RegisterInterfaceTypes(registry cdctypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (rc randomClient) RequestRandom(request RequestRandomRequest, basTx sdk.BaseTx) (RequestRandomResp, ctypes.ResultTx, error) {
	author, err := rc.QueryAddress(basTx.From, basTx.Password)
	if err != nil {
		return RequestRandomResp{}, ctypes.ResultTx{}, nil
	}

	msg := &MsgRequestRandom{
		BlockInterval: request.BlockInterval,
		Consumer:      author.String(),
		Oracle:        request.Oracle,
		ServiceFeeCap: request.ServiceFeeCap,
	}
	result, err := rc.BuildAndSend([]sdk.Msg{msg}, basTx)
	if err != nil {
		return RequestRandomResp{}, ctypes.ResultTx{}, err
	}
	reqID, e := sdk.StringifyEvents(result.TxResult.Events).GetValue(eventTypeRequestRequestRandom, attributeKeyRequestID)
	if e != nil {
		return RequestRandomResp{}, result, sdkerrors.Wrapf(ErrEventsGetValue, e.Error())
	}
	generateHeight, e := sdk.StringifyEvents(result.TxResult.Events).GetValue(eventTypeRequestRequestRandom, attributeKeyGenerateHeight)
	if e != nil {
		return RequestRandomResp{}, result, sdkerrors.Wrapf(ErrEventsGetValue, e.Error())
	}
	height, e := strconv.Atoi(generateHeight)
	if e != nil {
		return RequestRandomResp{}, result, sdkerrors.Wrapf(ErrAtoi, e.Error())
	}

	res := RequestRandomResp{
		Height: int64(height),
		ReqID:  reqID,
	}
	return res, result, nil
}

func (rc randomClient) QueryRandom(reqID string) (QueryRandomResp, error) {
	if len(reqID) == 0 {
		return QueryRandomResp{}, sdkerrors.Wrapf(ErrInvalidReqID, "reqId is required")
	}

	conn, err := rc.GenConn()
	if err != nil {
		return QueryRandomResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Random(
		context.Background(),
		&QueryRandomRequest{ReqId: reqID},
	)
	if err != nil {
		return QueryRandomResp{}, sdkerrors.Wrapf(ErrQueryRandom, err.Error())
	}
	return res.Random.Convert().(QueryRandomResp), nil
}

func (rc randomClient) QueryRandomRequestQueue(height int64) ([]QueryRandomRequestQueueResp, error) {
	if height == 0 {
		return []QueryRandomRequestQueueResp{}, nil
	}

	conn, err := rc.GenConn()
	if err != nil {
		return []QueryRandomRequestQueueResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}
	res, err := NewQueryClient(conn).RandomRequestQueue(
		context.Background(),
		&QueryRandomRequestQueueRequest{Height: height},
	)
	if err != nil {
		return []QueryRandomRequestQueueResp{}, sdkerrors.Wrapf(ErrQueryRequestQueue, err.Error())
	}
	return Requests(res.Requests).Convert().([]QueryRandomRequestQueueResp), nil
}
