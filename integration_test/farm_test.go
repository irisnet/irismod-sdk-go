package integrationtest

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/stretchr/testify/require"

	"github.com/irisnet/irismod-sdk-go/farm"
)

func (s IntegrationTestSuite) TestFarm() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      200000,
		Memo:     "test",
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}
	Request := farm.CreatePoolRequest{
		Name:         "farm_pool",
		Description:  "farm_test",
		LpTokenDenom: "",
		StartHeight:  0,
	}

	result, err := s.Farm.CreatePool(Request, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result.Hash)

	AdjustPoolRequest := farm.AdjustPoolRequest{
		PoolName: "test_farm_pool",
	}

	result, err = s.Farm.AdjustPool(AdjustPoolRequest, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result.Hash)

	StakeRequest := farm.StakeRequest{
		PoolName: "test_farm_pool",
	}

	result, err = s.Farm.Stake(StakeRequest, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result.Hash)

	HarvestRequest := farm.HarvestRequest{
		PoolName: "test_farm_pool",
	}

	result, err = s.Farm.Harvest(HarvestRequest, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result.Hash)

	UnstakeRequest := farm.UnstakeRequest{
		PoolName: "test_farm_pool",
	}

	result, err = s.Farm.Unstake(UnstakeRequest, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result.Hash)

	DestroyPoolRequest := farm.DestroyPoolRequest{
		PoolName: "test_farm_pool",
	}

	result, err = s.Farm.DestroyPool(DestroyPoolRequest, baseTx)
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), result.Hash)

}
