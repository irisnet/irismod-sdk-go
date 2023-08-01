package nft

import (
	"github.com/cosmos/cosmos-sdk/types/query"
	sdk "github.com/irisnet/core-sdk-go/types"
)

// expose NFT module api for user
type Client interface {
	sdk.Module

	IssueDenom(request IssueDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	MintNFT(request MintNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	EditNFT(request EditNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferNFT(request TransferNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	BurnNFT(request BurnNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)
	TransferDenom(request TransferDenomRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error)

	QuerySupply(denomID, creator string) (uint64, sdk.Error)
	QueryOwner(creator, denomID string, pagination *query.PageRequest) (QueryOwnerResp, sdk.Error)
	QueryCollection(denomID string, pagination *query.PageRequest) (QueryCollectionResp, sdk.Error)
	QueryDenom(denomID string) (QueryDenomResp, sdk.Error)
	QueryDenoms(pagination *query.PageRequest) ([]QueryDenomResp, sdk.Error)
	QueryNFT(denomID, tokenID string) (QueryNFTResp, sdk.Error)
}

type IssueDenomRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Schema      string `json:"schema"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
	Uri         string `json:"uri"`
	UriHash     string `json:"uri_hash"`
	Data        string `json:"data"`
}

type MintNFTRequest struct {
	Denom     string `json:"denom"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Recipient string `json:"recipient"`
	URIHash   string `json:"uri_hash"`
}

type EditNFTRequest struct {
	Denom   string `json:"denom"`
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	URIHash string `json:"uri_hash"`
}

type TransferNFTRequest struct {
	Denom     string `json:"denom"`
	ID        string `json:"id"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Name      string `json:"name"`
	Recipient string `json:"recipient"`
	URIHash   string `json:"uri_hash"`
}

type BurnNFTRequest struct {
	Denom string `json:"denom"`
	ID    string `json:"id"`
}

type TransferDenomRequest struct {
	ID        string `json:"id"`
	Recipient string `json:"recipient"`
}

type QueryOwnerResp struct {
	OwnerResp  *OwnerResp    `json:"owner"`
	Pagination *PageResponse `json:"pagination"`
}

type OwnerResp struct {
	Address string `json:"address" yaml:"address"`
	IDCs    []IDC  `json:"idcs" yaml:"idcs"`
}

// IDC defines a set of nft ids that belong to a specific
type IDC struct {
	Denom    string   `json:"denom" yaml:"denom"`
	TokenIDs []string `json:"token_ids" yaml:"token_ids"`
}

type PageResponse struct {
	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	NextKey []byte `json:"next_key"`
	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total uint64 `json:"total"`
}

// BaseNFT non fungible token definition
type QueryNFTResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	Creator string `json:"creator"`
	URIHash string `json:"uri_hash"`
}

type QueryDenomResp struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Schema      string `json:"schema"`
	Creator     string `json:"creator"`
	Description string `json:"description"`
	Uri         string `json:"uri"`
	UriHash     string `json:"uri_hash"`
	Data        string `json:"data"`
}

type QueryCollectionResp struct {
	Denom QueryDenomResp `json:"denom" yaml:"denom"`
	NFTs  []QueryNFTResp `json:"nfts" yaml:"nfts"`
}
