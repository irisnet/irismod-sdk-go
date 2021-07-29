package farm

import (
	"context"
	"github.com/irisnet/core-sdk-go/common/codec"
	"github.com/irisnet/core-sdk-go/common/codec/types"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type farmClient struct {
	sdk.BaseClient
	codec.Marshaler
}

func NewClient(bc sdk.BaseClient, cdc codec.Marshaler) Client {
	return farmClient{
		BaseClient: bc,
		Marshaler:  cdc,
	}
}

func (fc farmClient) Name() string {
	return ModuleName
}

func (fc farmClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (fc farmClient) CreatePool(request CreatePoolRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	author, err := fc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgCreatePool{
		Name:           request.Name,
		Description:    request.Description,
		LpTokenDenom:   request.LpTokenDenom,
		StartHeight:    request.StartHeight,
		RewardPerBlock: request.RewardPerBlock,
		TotalReward:    request.TotalReward,
		Editable:       request.Editable,
		Creator:        author.String(),
	}
	return fc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (fc farmClient) DestroyPool(request DestroyPoolRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	author, err := fc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgDestroyPool{
		PoolName: request.PoolName,
		Creator:  author.String(),
	}
	return fc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (fc farmClient) AdjustPool(request AdjustPoolRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	author, err := fc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgAdjustPool{
		PoolName:         request.PoolName,
		AdditionalReward: request.AdditionalReward,
		RewardPerBlock:   request.RewardPerBlock,
		Creator:          author.String(),
	}
	return fc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (fc farmClient) Stake(request StakeRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	author, err := fc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgStake{
		PoolName: request.PoolName,
		Amount:   request.Amount,
		Sender:   author.String(),
	}
	return fc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (fc farmClient) Unstake(request UnstakeRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	author, err := fc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgUnstake{
		PoolName: request.PoolName,
		Amount:   request.Amount,
		Sender:   author.String(),
	}
	return fc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (fc farmClient) Harvest(request HarvestRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	author, err := fc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgUnstake{
		PoolName: request.PoolName,
		Sender:   author.String(),
	}
	return fc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (fc farmClient) QueryFarmPools(request QueryFarmPoolsRequest) (QueryFarmPoolsResponse, sdk.Error) {
	conn, err := fc.GenConn()
	if err != nil {
		return QueryFarmPoolsResponse{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).FarmPools(
		context.Background(),
		&QueryFarmPoolsRequest{
			Pagination: request.Pagination,
		},
	)
	if err != nil {
		return QueryFarmPoolsResponse{}, sdk.Wrap(err)
	}

	return QueryFarmPoolsResponse{
		Pools:      res.Pools,
		Pagination: res.Pagination,
	}, nil
}

func (fc farmClient) QueryFarmPool(request QueryFarmPoolRequest) (QueryFarmPoolResponse, sdk.Error) {
	conn, err := fc.GenConn()
	if err != nil {
		return QueryFarmPoolResponse{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).FarmPool(
		context.Background(),
		&QueryFarmPoolRequest{
			Name: request.Name,
		},
	)
	if err != nil {
		return QueryFarmPoolResponse{}, sdk.Wrap(err)
	}

	return QueryFarmPoolResponse{
		Pool: res.Pool,
	}, nil
}

func (fc farmClient) QueryFarmer(request QueryFarmerRequest) (QueryFarmerResponse, sdk.Error) {
	conn, err := fc.GenConn()
	if err != nil {
		return QueryFarmerResponse{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Farmer(
		context.Background(),
		&QueryFarmerRequest{
			Farmer:   request.Farmer,
			PoolName: request.PoolName,
		},
	)
	if err != nil {
		return QueryFarmerResponse{}, sdk.Wrap(err)
	}

	return QueryFarmerResponse{
		List:   res.List,
		Height: res.Height,
	}, nil
}

func (fc farmClient) QueryParams(request QueryParamsRequest) (QueryParamsResponse, sdk.Error) {
	conn, err := fc.GenConn()
	if err != nil {
		return QueryParamsResponse{}, sdk.Wrap(err)
	}

	res, err := NewQueryClient(conn).Params(
		context.Background(),
		&QueryParamsRequest{},
	)
	if err != nil {
		return QueryParamsResponse{}, sdk.Wrap(err)
	}

	return QueryParamsResponse{
		Params: res.Params,
	}, nil
}
