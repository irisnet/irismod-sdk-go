package nft

import (
	sdk "github.com/irisnet/core-sdk-go/types"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"
)

// expose NFT module api for user
type Client interface {
	sdk.Module

	IssueDenom(request IssueDenomRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	MintNFT(request MintNFTRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	EditNFT(request EditNFTRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	TransferNFT(request TransferNFTRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error)
	BurnNFT(request BurnNFTRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error)

	QuerySupply(denomID, creator string) (uint64, error)
	QueryOwner(creator, denomID string) (QueryOwnerResp, error)
	QueryCollection(denomID string) (QueryCollectionResp, error)
	QueryDenom(denomID string) (QueryDenomResp, error)
	QueryDenoms() ([]QueryDenomResp, error)
	QueryNFT(denomID, tokenID string) (QueryNFTResp, error)
}

type IssueDenomRequest struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type MintNFTRequest struct {
	Denom     string `json:"denom"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Recipient string `json:"recipient"`
}

type EditNFTRequest struct {
	Denom string `json:"denom"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	URI   string `json:"uri"`
	Data  string `json:"data"`
}

type TransferNFTRequest struct {
	Denom     string `json:"denom"`
	ID        string `json:"id"`
	URI       string `json:"uri"`
	Data      string `json:"data"`
	Name      string `json:"name"`
	Recipient string `json:"recipient"`
}

type BurnNFTRequest struct {
	Denom string `json:"denom"`
	ID    string `json:"id"`
}

// IDC defines a set of nft ids that belong to a specific
type IDC struct {
	Denom    string   `json:"denom" yaml:"denom"`
	TokenIDs []string `json:"token_ids" yaml:"token_ids"`
}

type QueryOwnerResp struct {
	Address string `json:"address" yaml:"address"`
	IDCs    []IDC  `json:"idcs" yaml:"idcs"`
}

// BaseNFT non fungible token definition
type QueryNFTResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URI     string `json:"uri"`
	Data    string `json:"data"`
	Creator string `json:"creator"`
}

type QueryDenomResp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Schema  string `json:"schema"`
	Creator string `json:"creator"`
}

type QueryCollectionResp struct {
	Denom QueryDenomResp `json:"denom" yaml:"denom"`
	NFTs  []QueryNFTResp `json:"nfts" yaml:"nfts"`
}
