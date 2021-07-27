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

func (fc farmClient) CreatePool(request CreatePoolRequest) sdk.Error {

	return nil
}

func (fc farmClient) DestroyPool(request DestroyPoolRequest) sdk.Error {

	return nil
}

func (fc farmClient) AdjustPool(request AdjustPoolRequest) sdk.Error {

	return nil
}

func (fc farmClient) Stake(request StakeRequest) sdk.Error {

	return nil
}

func (fc farmClient) Unstake(request UnstakeRequest) sdk.Error {

	return nil
}

func (fc farmClient) Harvest(request HarvestRequest) sdk.Error {

	return nil
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
