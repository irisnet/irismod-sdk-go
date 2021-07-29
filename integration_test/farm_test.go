package integrationtest

import sdk "github.com/irisnet/core-sdk-go/types"

func (s IntegrationTestSuite) TestFarm() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      200000,
		Memo:     "test",
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}

}
