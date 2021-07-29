package farm

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/core-sdk-go/types/query"
)

// expose Farm module api for user
type Client interface {
	sdk.Module

	CreatePool(request CreatePoolRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	DestroyPool(request DestroyPoolRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	AdjustPool(request AdjustPoolRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	Stake(request StakeRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	Unstake(request UnstakeRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	Harvest(request HarvestRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)

	QueryFarmPools(request QueryFarmPoolsRequest) (QueryFarmPoolsResponse, sdk.Error)
	QueryFarmPool(request QueryFarmPoolRequest) (QueryFarmPoolResponse, sdk.Error)
	QueryFarmer(request QueryFarmerRequest) (QueryFarmerResponse, sdk.Error)
	QueryParams(request QueryParamsRequest) (QueryParamsResponse, sdk.Error)
}

type CreatePoolRequest struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	LpTokenDenom   string    `json:"lp_token_denom"`
	StartHeight    int64     `json:"start_height"`
	RewardPerBlock sdk.Coins `json:"reward_per_block"`
	TotalReward    sdk.Coins `json:"total_reward"`
	Editable       bool      `json:"editable"`
	Creator        string    `json:"creator"`
}

type DestroyPoolRequest struct {
	PoolName string `json:"pool_name"`
	Creator  string `json:"creator"`
}

type AdjustPoolRequest struct {
	PoolName         string    `json:"pool_name"`
	AdditionalReward sdk.Coins `json:"additional_reward"`
	RewardPerBlock   sdk.Coins `json:"reward_per_block"`
	Creator          string    `json:"creator"`
}

type StakeRequest struct {
	PoolName string   `json:"pool_name"`
	Amount   sdk.Coin `json:"amount"`
	Sender   string   `json:"sender"`
}

type UnstakeRequest struct {
	PoolName string   `json:"pool_name"`
	Amount   sdk.Coin `json:"amount"`
	Sender   string   `json:"sender"`
}

type HarvestRequest struct {
	PoolName string `json:"pool_name"`
	Sender   string `json:"sender"`
}

type QueryFarmPoolsResp struct {
	Pools      []*FarmPoolEntry    `json:"pools"`
	Pagination *query.PageResponse `json:"pagination"`
}
