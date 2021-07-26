package farm

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

// expose Gov module api for user
type Client interface {
	sdk.Module

	CreatePool(request CreatePoolRequest) sdk.Error
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
	Title string `json:"title"`
}

type DestroyPoolRequest struct {
	Title string `json:"title"`
}

type AdjustPoolRequest struct {
	Title string `json:"title"`
}

type StakeRequest struct {
	Title string `json:"title"`
}

type UnstakeRequest struct {
	Title string `json:"title"`
}

type HarvestRequest struct {
	Title string `json:"title"`
}
