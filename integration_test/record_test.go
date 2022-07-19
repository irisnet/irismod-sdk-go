package integrationtest

import (
	"fmt"
	"github.com/irisnet/irismod-sdk-go/record"
	"github.com/stretchr/testify/require"

	sdk "github.com/irisnet/core-sdk-go/types"
)

func (s IntegrationTestSuite) TestRecord() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      200000,
		Memo:     "test",
		Mode:     sdk.Sync,
		Password: s.Account().Password,
	}

	num := 5
	contents := make([]record.Content, num)
	for i := 0; i < num; i++ {
		contents[i] = record.Content{
			Digest:     s.RandStringOfLength(10),
			DigestAlgo: s.RandStringOfLength(5),
			URI:        fmt.Sprintf("https://%s", s.RandStringOfLength(10)),
			Meta:       s.RandStringOfLength(20),
		}
	}

	req := record.CreateRecordRequest{
		Contents: contents,
	}

	resp, err := s.Record.CreateRecord(req, baseTx)
	if err != nil {
		panic(err)
	}
	require.NoError(s.T(), err)
	require.NotEmpty(s.T(), resp.Hash)

	if resp.RecordId != "" {
		request := record.QueryRecordReq{
			RecordID: resp.RecordId,
			Prove:    true,
			Height:   resp.Height,
		}

		result, err := s.Record.QueryRecord(request)
		if err != nil {
			panic(err)
		}
		require.NoError(s.T(), err)
		require.NotEmpty(s.T(), result.Record.Contents)

		for i := 0; i < num; i++ {
			require.EqualValues(s.T(), contents[i], result.Record.Contents[i])
		}
	}
}
