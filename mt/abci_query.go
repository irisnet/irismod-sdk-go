package mt

import (
	"context"
	"errors"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
)

const (
	FunctionPathQueryDenom  = "/irismod.mt.Query/Denom"
	FunctionPathQueryDenoms = "/irismod.mt.Query/Denoms"
	FunctionPathQueryMT     = "/irismod.mt.Query/MT"
)

func (cli *Client) ABCIQueryClass(classId string, height int64) (*QueryClassResp, error) {
	if len(classId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidClassId, "class id is required")
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
		FunctionPathQueryDenom, reqBz, opts)
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
		ID:    grpcResp.Denom.Id,
		Name:  grpcResp.Denom.Name,
		Data:  grpcResp.Denom.Data,
		Owner: grpcResp.Denom.Owner,
	}

	return &resp, nil
}

func (cli *Client) ABCIQueryClass2(classId string, height int64) (*QueryClassResp, error) {
	var resp QueryClassResp

	if len(classId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidClassId, "class id is required")
	}

	if height < 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidHeight, "height must be not less than 0")
	}

	grpcReq := &QueryDenomsRequest{}

	reqBz, err := cli.Marshaler().Marshal(grpcReq)
	if err != nil {
		return nil, err
	}

	opts := rpcclient.ABCIQueryOptions{
		Height: height,
		Prove:  false,
	}

	result, err := cli.ABCIQueryWithOptions(context.Background(),
		FunctionPathQueryDenoms, reqBz, opts)
	if err != nil {
		return nil, err
	}

	if !result.Response.IsOK() {
		return nil, errors.New(fmt.Sprint(result.Response.Log))
	}

	var grpcResp QueryDenomsResponse
	err = cli.Marshaler().Unmarshal(result.Response.Value, &grpcResp)
	if err != nil {
		return nil, err
	}

	for _, denom := range grpcResp.Denoms {
		if denom.Id == classId {
			resp = QueryClassResp{
				ID:    denom.Id,
				Name:  denom.Name,
				Data:  denom.Data,
				Owner: denom.Owner,
			}
			break
		}
	}

	return &resp, nil
}

func (cli *Client) ABCIQueryMT(classId, tokenId string, height int64) (*QueryMTResp, error) {
	if len(classId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidClassId, "class id is required")
	}

	if len(tokenId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidTokenID, "token id is required")
	}

	if height < 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidHeight, "height must be not less than 0")
	}

	grpcReq := &QueryMTRequest{
		DenomId: classId,
		MtId:    tokenId,
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
		FunctionPathQueryMT, reqBz, opts)
	if err != nil {
		return nil, err
	}

	if !result.Response.IsOK() {
		return nil, errors.New(fmt.Sprint(result.Response.Log))
	}

	var grpcResp QueryMTResponse
	err = cli.Marshaler().Unmarshal(result.Response.Value, &grpcResp)
	if err != nil {
		return nil, err
	}

	resp := QueryMTResp{
		ID:     grpcResp.Mt.Id,
		Supply: grpcResp.Mt.Supply,
		Data:   grpcResp.Mt.Data,
	}

	return &resp, nil
}
