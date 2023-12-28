package mt

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type IABIClient interface {
	ABCIQueryClass(classId string, height int64) (*QueryClassResp, error)
	ABCIQueryClass2(classId string, height int64) (*QueryClassResp, error)
	ABCIQueryMT(classId, tokenId string, height int64) (*QueryMTResp, error)
}

// IClient expose NFT module api for user
type IClient interface {
	sdk.Module
	IABIClient
	CreateClass(request IssueClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferClass(request TransferClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	MintMT(request MintMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	AddMT(request AddMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	EditMT(request EditMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferMT(request TransferMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	BurnMT(request BurnMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)

	QuerySupply(ClassID, creator string) (uint64, sdk.Error)
	QueryClasses(pageReq *query.PageRequest) (*QueryClassesResp, sdk.Error)
	QueryClass(ClassID string) (QueryClassResp, sdk.Error)
	QueryMTSupply(ClassID, mtID string) (uint64, sdk.Error)
	QueryMTs(ClassID string, pageReq *query.PageRequest) (*QueryMtsResp, sdk.Error)
	QueryMT(ClassID, mtID string) (QueryMTResp, sdk.Error)
	QueryBalances(ClassID, owner string, pagination *query.PageRequest) (*QueryBalancesResp, sdk.Error)
}
