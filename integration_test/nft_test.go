package integration_test

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/irismod-sdk-go/nft"
)

const (
	testClassId               = "testclassid"
	testClassName             = "testclassname"
	testClassSchema           = "testclassschema"
	testClassSymbol           = "testclasssymbol"
	testClassDescription      = "testclassdescription"
	testClassURI              = "http://cat.market/"
	testClassURIHash          = "testclassurihash"
	testClassData             = "{\"name\":\"cat\"}"
	testClassMintRestricted   = true
	testClassUpdateRestricted = false
)

func (s *ClientTestSuite) TestNFT() {
	baseTx := sdk.BaseTx{
		From:     s.Account().Name,
		Gas:      400000,
		Memo:     "",
		Mode:     sdk.Commit,
		Password: s.Account().Password,
	}
	req := nft.NewCreateClassRequest(
		testClassId,
		testClassName,
		testClassSchema,
		testClassSymbol,
		testClassDescription,
		testClassURI,
		testClassURIHash,
		testClassData,
		testClassMintRestricted,
		testClassUpdateRestricted,
	)
	resp, err := s.NFTClient.CreateClass(req, baseTx)
	s.Require().NoError(err)
	s.T().Log(resp)
}
