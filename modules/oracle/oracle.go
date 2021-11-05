package oracle

import (
	"context"

	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type oracleClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(baseClient sdk.BaseClient, marshaler codec.Codec) Client {
	return oracleClient{
		BaseClient: baseClient,
		Codec:      marshaler,
	}
}

func (oc oracleClient) Name() string {
	return ModuleName
}

func (oc oracleClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (oc oracleClient) CreateFeed(request CreateFeedRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	serviceFeeCap, e := oc.ToMinCoin(request.ServiceFeeCap...)
	if e != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrToMintCoin, err.Error())
	}

	msg := &MsgCreateFeed{
		FeedName:          request.FeedName,
		LatestHistory:     request.LatestHistory,
		Description:       request.Description,
		Creator:           sender.String(),
		ServiceName:       request.ServiceName,
		Providers:         request.Providers,
		Input:             request.Input,
		Timeout:           request.Timeout,
		ServiceFeeCap:     serviceFeeCap,
		RepeatedFrequency: request.RepeatedFrequency,
		AggregateFunc:     request.AggregateFunc,
		ValueJsonPath:     request.ValueJsonPath,
		ResponseThreshold: request.ResponseThreshold,
	}
	return oc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (oc oracleClient) StartFeed(feedName string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	msg := &MsgStartFeed{
		FeedName: feedName,
		Creator:  sender.String(),
	}
	return oc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (oc oracleClient) PauseFeed(feedName string, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	msg := &MsgPauseFeed{
		FeedName: feedName,
		Creator:  sender.String(),
	}
	return oc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (oc oracleClient) EditFeed(request EditFeedRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := oc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	serviceFeeCap, e := oc.ToMinCoin(request.ServiceFeeCap...)
	if e != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrToMintCoin, err.Error())
	}

	msg := &MsgEditFeed{
		FeedName:          request.FeedName,
		Description:       request.Description,
		LatestHistory:     request.LatestHistory,
		Providers:         request.Providers,
		Timeout:           request.Timeout,
		ServiceFeeCap:     serviceFeeCap,
		RepeatedFrequency: request.RepeatedFrequency,
		ResponseThreshold: request.ResponseThreshold,
		Creator:           sender.String(),
	}
	return oc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (oc oracleClient) QueryFeed(feedName string) (QueryFeedResp, error) {
	if len(feedName) == 0 {
		return QueryFeedResp{}, sdkerrors.Wrapf(ErrInvalidFeedName, "feedName is required")
	}

	conn, err := oc.GenConn()
	if err != nil {
		return QueryFeedResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Feed(
		context.Background(),
		&QueryFeedRequest{FeedName: feedName},
	)
	if err != nil {
		return QueryFeedResp{}, sdkerrors.Wrapf(ErrQueryFeed, err.Error())
	}
	return res.Feed.Convert().(QueryFeedResp), nil
}

func (oc oracleClient) QueryFeeds(state string) ([]QueryFeedResp, error) {
	// todo state (whether state is required)
	if len(state) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidState, "state is required")
	}

	conn, err := oc.GenConn()
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Feeds(
		context.Background(),
		&QueryFeedsRequest{State: state},
	)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrQueryFeed, err.Error())
	}
	return feedContexts(res.Feeds).Convert().([]QueryFeedResp), nil
}

func (oc oracleClient) QueryFeedValue(feedName string) ([]QueryFeedValueResp, error) {
	if len(feedName) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidFeedName, "feedName is required")
	}

	conn, err := oc.GenConn()
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).FeedValue(
		context.Background(),
		&QueryFeedValueRequest{FeedName: feedName},
	)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrQueryFeedValue, err.Error())
	}
	return feedValues(res.FeedValues).Convert().([]QueryFeedValueResp), nil
}
