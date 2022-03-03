package mt

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type Client interface {
	sdk.Module

	IssueDenom(request IssueDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	MintMT(request MintMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	EditMT(request EditMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferMT(request TransferMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	BurnMT(request BurnMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferDenom(request TransferDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)

	QuerySupply(denomID, creator string) (uint64, sdk.Error)
	QueryDenoms() ([]QueryDenomResp, sdk.Error)
	QueryDenom(denomID string) (QueryDenomResp, sdk.Error)
	QueryMTSupply(denomID, mtID string) (uint64, sdk.Error)
	QueryMTs(denomID, mtID string, pagination *query.PageRequest) ([]QueryMTResp, sdk.Error)
	QueryMT(denomID, mtID string) (QueryMTResp, sdk.Error)
	QueryBalances(denomID, creator string, pagination *query.PageRequest) ([]Balance, sdk.Error)
}

type IssueDenomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Data []byte `json:"data"`
}

type MintMTRequest struct {
	ID        string `json:"id"`
	Denom     string `json:"denom"`
	Amount    uint64 `json:"amount"`
	Data      []byte `json:"data"`
	Recipient string `json:"recipient"`
}

type EditMTRequest struct {
	ID    string `json:"id"`
	Denom string `json:"denom"`
	Data  []byte `json:"data"`
}

type TransferMTRequest struct {
	ID        string `json:"id"`
	Denom     string `json:"denom"`
	Amount    uint64 `json:"amount"`
	Recipient string `json:"recipient"`
}

type BurnMTRequest struct {
	ID     string `json:"id"`
	Denom  string `json:"denom"`
	Amount uint64 `json:"amount"`
}

type TransferDenomRequest struct {
	ID        string `json:"id"`
	Recipient string `json:"recipient"`
}

// IDC defines a set of mt ids that belong to a specific
type IDC struct {
	Denom    string   `json:"denom" yaml:"denom"`
	TokenIDs []string `json:"token_ids" yaml:"token_ids"`
}

type QueryOwnerResp struct {
	Address string `json:"address" yaml:"address"`
	IDCs    []IDC  `json:"idcs" yaml:"idcs"`
}

// BaseMT non fungible token definition
type QueryMTResp struct {
	ID     string `json:"id"`
	Supply uint64 `json:"supply"`
	Data   []byte `json:"data"`
}

type QueryDenomResp struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Data  []byte `json:"data"`
	Owner string `json:"owner"`
}

type QueryBalanceResp struct {
	Balance    []Balance           `json:"balance"`
	Pagination *query.PageResponse `json:"pagination"`
}

type QueryCollectionResp struct {
	Denom QueryDenomResp `json:"denom" yaml:"denom"`
	MTs   []QueryMTResp  `json:"mts" yaml:"mts"`
}
