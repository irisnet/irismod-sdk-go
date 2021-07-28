package farm

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

// expose Farm module api for user
type Client interface {
	sdk.Module

	CreatePool(request CreatePoolRequest) (MsgCreatePoolResponse, sdk.Error)
	DestroyPool(request DestroyPoolRequest) sdk.Error
	AdjustPool(request AdjustPoolRequest) sdk.Error
	Stake(request StakeRequest) sdk.Error
	Unstake(request UnstakeRequest) sdk.Error
	Harvest(request HarvestRequest) sdk.Error

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
