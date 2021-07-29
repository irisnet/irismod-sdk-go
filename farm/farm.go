package farm

import (
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

	return QueryFarmPoolsResponse{}, nil
}

func (fc farmClient) QueryFarmPool(request QueryFarmPoolRequest) (QueryFarmPoolResponse, sdk.Error) {

	return QueryFarmPoolResponse{}, nil
}

func (fc farmClient) QueryFarmer(request QueryFarmerRequest) (QueryFarmerResponse, sdk.Error) {

	return QueryFarmerResponse{}, nil
}

func (fc farmClient) QueryParams(request QueryParamsRequest) (QueryParamsResponse, sdk.Error) {

	return QueryParamsResponse{}, nil
}
