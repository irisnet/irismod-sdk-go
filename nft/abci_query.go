package nft

import (
	"context"
	"errors"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

func (cli *Client) ABCIQueryClass(classId string, height int64) (*QueryClassResp, error) {
	if len(classId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidDenom, "class id is required")
	}

	if height < 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidHeight, "height must be not less than 0")
	}

	grpcReq := &QueryDenomRequest{
		DenomId: classId,
	}

	reqBz, err := cli.Marshaler().Marshal(grpcReq)
	if err != nil {
		return nil, err
	}

	opts := rpcclient.ABCIQueryOptions{
		Height: height,
		Prove:  false,
	}

	result, err := cli.ABCIQueryWithOptions(context.Background(),
		"/irismod.nft.Query/Denom", reqBz, opts)
	if err != nil {
		return nil, err
	}

	if !result.Response.IsOK() {
		return nil, errors.New(fmt.Sprint(result.Response.Log))
	}

	var grpcResp QueryDenomResponse
	err = cli.Marshaler().Unmarshal(result.Response.Value, &grpcResp)
	if err != nil {
		return nil, err
	}

	resp := QueryClassResp{
		ID:               grpcResp.Denom.Id,
		Name:             grpcResp.Denom.Name,
		Schema:           grpcResp.Denom.Schema,
		Symbol:           grpcResp.Denom.Symbol,
		Creator:          grpcResp.Denom.Creator,
		Description:      grpcResp.Denom.Description,
		Uri:              grpcResp.Denom.Uri,
		UriHash:          grpcResp.Denom.UriHash,
		Data:             grpcResp.Denom.Data,
		MintRestricted:   grpcResp.Denom.MintRestricted,
		UpdateRestricted: grpcResp.Denom.UpdateRestricted,
	}

	return &resp, nil
}

func (cli *Client) ABCIQueryNFT(classId, tokenId string, height int64) (*QueryNFTResp, error) {
	if len(classId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidDenom, "class id is required")
	}

	if len(tokenId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidTokenID, "token id is required")
	}

	if height < 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidHeight, "height must be not less than 0")
	}

	grpcReq := &QueryNFTRequest{
		DenomId: classId,
		TokenId: tokenId,
	}

	reqBz, err := cli.Marshaler().Marshal(grpcReq)
	if err != nil {
		return nil, err
	}

	opts := rpcclient.ABCIQueryOptions{
		Height: height,
		Prove:  false,
	}

	result, err := cli.ABCIQueryWithOptions(context.Background(),
		"/irismod.nft.Query/NFT", reqBz, opts)
	if err != nil {
		return nil, err
	}

	if !result.Response.IsOK() {
		return nil, errors.New(fmt.Sprint(result.Response.Log))
	}

	var grpcResp QueryNFTResponse
	err = cli.Marshaler().Unmarshal(result.Response.Value, &grpcResp)
	if err != nil {
		return nil, err
	}

	resp := QueryNFTResp{
		ID:      grpcResp.NFT.Id,
		Name:    grpcResp.NFT.Name,
		URI:     grpcResp.NFT.URI,
		Data:    grpcResp.NFT.Data,
		Owner:   grpcResp.NFT.Owner,
		URIHash: grpcResp.NFT.UriHash,
	}

	return &resp, nil
}
