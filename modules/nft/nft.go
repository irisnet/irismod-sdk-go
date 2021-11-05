package nft

import (
	"context"

	sdkerrors "github.com/irisnet/core-sdk-go/types/errors"
	ctypes "github.com/tendermint/tendermint/rpc/core/types"

	"github.com/irisnet/core-sdk-go/codec"
	"github.com/irisnet/core-sdk-go/codec/types"

	sdk "github.com/irisnet/core-sdk-go/types"
)

type nftClient struct {
	sdk.BaseClient
	codec.Codec
}

func NewClient(bc sdk.BaseClient, cdc codec.Codec) Client {
	return nftClient{
		BaseClient: bc,
		Codec:      cdc,
	}
}

func (nc nftClient) Name() string {
	return ModuleName
}

func (nc nftClient) RegisterInterfaceTypes(registry types.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (nc nftClient) IssueDenom(request IssueDenomRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	msg := &MsgIssueDenom{
		Id:     request.ID,
		Name:   request.Name,
		Schema: request.Schema,
		Sender: sender.String(),
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) MintNFT(request MintNFTRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	var recipient = sender.String()
	if len(request.Recipient) > 0 {
		if err := sdk.ValidateAccAddress(request.Recipient); err != nil {
			return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrValidateAccAddress, err.Error())
		}
		recipient = request.Recipient
	}

	msg := &MsgMintNFT{
		Id:        request.ID,
		DenomId:   request.Denom,
		Name:      request.Name,
		URI:       request.URI,
		Data:      request.Data,
		Sender:    sender.String(),
		Recipient: recipient,
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) EditNFT(request EditNFTRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	msg := &MsgEditNFT{
		Id:      request.ID,
		Name:    request.Name,
		DenomId: request.Denom,
		URI:     request.URI,
		Data:    request.Data,
		Sender:  sender.String(),
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) TransferNFT(request TransferNFTRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	if err := sdk.ValidateAccAddress(request.Recipient); err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrValidateAccAddress, err.Error())
	}

	msg := &MsgTransferNFT{
		Id:        request.ID,
		Name:      request.Name,
		DenomId:   request.Denom,
		URI:       request.URI,
		Data:      request.Data,
		Sender:    sender.String(),
		Recipient: request.Recipient,
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) BurnNFT(request BurnNFTRequest, baseTx sdk.BaseTx) (ctypes.ResultTx, error) {
	sender, err := nc.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return ctypes.ResultTx{}, sdkerrors.Wrapf(ErrQueryAddress, err.Error())
	}

	msg := &MsgBurnNFT{
		Sender:  sender.String(),
		Id:      request.ID,
		DenomId: request.Denom,
	}
	return nc.BuildAndSend([]sdk.Msg{msg}, baseTx)
}

func (nc nftClient) QuerySupply(denom, creator string) (uint64, error) {
	if len(denom) == 0 {
		return 0, sdkerrors.Wrapf(ErrInvalidDenom, "denom is required")
	}

	if err := sdk.ValidateAccAddress(creator); err != nil {
		return 0, sdkerrors.Wrapf(ErrValidateAccAddress, err.Error())
	}

	conn, err := nc.GenConn()
	if err != nil {
		return 0, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Supply(
		context.Background(),
		&QuerySupplyRequest{
			Owner:   creator,
			DenomId: denom,
		},
	)
	if err != nil {
		return 0, sdkerrors.Wrapf(ErrQuerySupply, err.Error())
	}

	return res.Amount, nil
}

func (nc nftClient) QueryOwner(creator, denom string) (QueryOwnerResp, error) {
	if len(denom) == 0 {
		return QueryOwnerResp{}, sdkerrors.Wrapf(ErrInvalidDenom, "denom is required")
	}

	if err := sdk.ValidateAccAddress(creator); err != nil {
		return QueryOwnerResp{}, sdkerrors.Wrapf(ErrValidateAccAddress, err.Error())
	}

	conn, err := nc.GenConn()
	if err != nil {
		return QueryOwnerResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Owner(
		context.Background(),
		&QueryOwnerRequest{
			Owner:   creator,
			DenomId: denom,
		},
	)
	if err != nil {
		return QueryOwnerResp{}, sdkerrors.Wrapf(ErrQueryOwner, err.Error())
	}

	return res.Owner.Convert().(QueryOwnerResp), nil
}

func (nc nftClient) QueryCollection(denom string) (QueryCollectionResp, error) {
	if len(denom) == 0 {
		return QueryCollectionResp{}, sdkerrors.Wrapf(ErrInvalidDenom, "denom is required")
	}

	conn, err := nc.GenConn()
	if err != nil {
		return QueryCollectionResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Collection(
		context.Background(),
		&QueryCollectionRequest{DenomId: denom},
	)
	if err != nil {
		return QueryCollectionResp{}, sdkerrors.Wrapf(ErrQueryCollection, err.Error())
	}

	return res.Collection.Convert().(QueryCollectionResp), nil
}

func (nc nftClient) QueryDenoms() ([]QueryDenomResp, error) {
	conn, err := nc.GenConn()
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Denoms(
		context.Background(),
		&QueryDenomsRequest{},
	)
	if err != nil {
		return nil, sdkerrors.Wrapf(ErrQueryDenoms, err.Error())
	}

	return denoms(res.Denoms).Convert().([]QueryDenomResp), nil
}

func (nc nftClient) QueryDenom(denom string) (QueryDenomResp, error) {
	conn, err := nc.GenConn()
	if err != nil {
		return QueryDenomResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).Denom(
		context.Background(),
		&QueryDenomRequest{DenomId: denom},
	)
	if err != nil {
		return QueryDenomResp{}, sdkerrors.Wrapf(ErrQueryDenoms, err.Error())
	}

	return res.Denom.Convert().(QueryDenomResp), nil
}

func (nc nftClient) QueryNFT(denom, tokenID string) (QueryNFTResp, error) {
	if len(denom) == 0 {
		return QueryNFTResp{}, sdkerrors.Wrapf(ErrInvalidDenom, "denom is required")
	}

	if len(tokenID) == 0 {
		return QueryNFTResp{}, sdkerrors.Wrapf(ErrInvalidTokenID, "tokenID is required")
	}

	conn, err := nc.GenConn()
	if err != nil {
		return QueryNFTResp{}, sdkerrors.Wrapf(ErrGenConn, err.Error())
	}

	res, err := NewQueryClient(conn).NFT(
		context.Background(),
		&QueryNFTRequest{
			DenomId: denom,
			TokenId: tokenID,
		},
	)
	if err != nil {
		return QueryNFTResp{}, sdkerrors.Wrapf(ErrQueryNFT, err.Error())
	}

	return res.NFT.Convert().(QueryNFTResp), nil
}
