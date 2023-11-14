package nft

import (
	sdk "github.com/irisnet/core-sdk-go/types"
)

// IClient expose NFT module api for user
type IClient interface {
	sdk.Module

	CreateClass(request CreateClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	TransferClass(request TransferClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	MintNFT(request MintNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	EditNFT(request EditNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	TransferNFT(request TransferNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)
	BurnNFT(request BurnNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error)

	QuerySupply(classId, creator string) (uint64, error)
	QueryOwner(creator, classId string, pagination PaginationRequest) (*QueryOwnerResp, error)
	QueryCollection(classId string, pagination PaginationRequest) (*QueryCollectionResp, error)
	QueryClass(classId string) (*QueryClassResp, error)
	QueryClasses(pagination PaginationRequest) (*QueryClassesResp, error)
	QueryNFT(classId, tokenID string) (*QueryNFTResp, error)

	ABCIQueryDenom(classId string, height int64) (*QueryClassResp, error)
	ABCIQueryNFT(classId, tokenId string, height int64) (*QueryNFTResp, error)
}
