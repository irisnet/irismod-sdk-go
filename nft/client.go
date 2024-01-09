package nft

import (
	"context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	sdk "github.com/irisnet/core-sdk-go/types"
)

type Client struct {
	sdk.BaseClient

	queryCli QueryClient
}

func NewClient(bc sdk.BaseClient) (IClient, error) {
	return &Client{
		BaseClient: bc,
		queryCli:   NewQueryClient(bc.GrpcConn()),
	}, nil
}

func (cli *Client) Name() string {
	return ModuleName
}
func (cli *Client) RegisterInterfaceTypes(registry codectypes.InterfaceRegistry) {
	RegisterInterfaces(registry)
}

func (cli *Client) CreateClass(request CreateClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to query address")
	}
	msg := &MsgIssueDenom{
		Id:               request.ID,
		Name:             request.Name,
		Schema:           request.Schema,
		Sender:           sender.String(),
		Symbol:           request.Symbol,
		MintRestricted:   request.MintRestricted,
		UpdateRestricted: request.UpdateRestricted,
		Data:             request.Data,
		Uri:              request.Uri,
		UriHash:          request.UriHash,
	}

	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) MintNFT(request MintNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to query address")
	}
	msg, err1 := request.ToMsg()
	if err1 != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to convert to msg")
	}

	msg.Sender = sender.String()

	if len(msg.Recipient) == 0 {
		msg.Recipient = sender.String()
	}

	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) EditNFT(request EditNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to query address")
	}
	msg, err1 := request.ToMsg()
	if err1 != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to convert to msg")
	}

	msg.Sender = sender.String()

	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) TransferNFT(request TransferNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to query address")
	}
	msg, err1 := request.ToMsg()
	if err1 != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to convert to msg")
	}

	msg.Sender = sender.String()

	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) BurnNFT(request BurnNFTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to query address")
	}
	msg, err1 := request.ToMsg()
	if err1 != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to convert to msg")
	}

	msg.Sender = sender.String()

	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) TransferClass(request TransferClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to query address")
	}
	msg, err1 := request.ToMsg()
	if err1 != nil {
		return sdk.ResultTx{}, sdkerrors.Wrap(err, "failed to convert to msg")
	}

	msg.Sender = sender.String()

	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

// QuerySupply queries the total supply of a given classId or owner
// classId is required
// owner is optional
func (cli *Client) QuerySupply(classId, owner string) (uint64, error) {
	if len(classId) == 0 {

		return 0, sdkerrors.Wrapf(ErrInvalidDenom, "denom is required")
	}

	req := &QuerySupplyRequest{
		DenomId: classId,
	}

	if len(owner) > 0 {
		if _, err := types.AccAddressFromBech32(owner); err != nil {
			return 0, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
		}
		req.Owner = owner
	}

	res, err := cli.queryCli.Supply(
		context.Background(),
		req,
	)
	if err != nil {
		return 0, sdkerrors.Wrapf(err, "failed to query supply: %s", err)
	}

	return res.Amount, nil
}

// QueryOwner queries the NFTs of the specified owner.
// owner is required
// classId is optional
func (cli *Client) QueryOwner(owner, classId string, pagination PaginationRequest) (*QueryOwnerResp, error) {
	if len(owner) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", owner)
	}

	req := &QueryOwnerRequest{
		Owner: owner,
	}

	if len(classId) > 0 {
		req.DenomId = classId
	}

	page := &query.PageRequest{
		CountTotal: false,
		Offset:     pagination.Offset,
	}

	if pagination.Limit > 0 {
		page.Limit = pagination.Limit
	} else {
		page.Limit = 100
	}

	if len(pagination.NextKey) > 0 {
		page.Key = []byte(pagination.NextKey)
	}

	req.Pagination = page

	result, err := cli.queryCli.Owner(
		context.Background(),
		req)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to query owner: %s", err)
	}

	idcs := make([]IDC, 0, len(result.Owner.IDCollections))

	for index, value := range result.Owner.IDCollections {
		idc := IDC{
			Class:    value.DenomId,
			TokenIDs: value.TokenIds,
		}
		idcs[index] = idc
	}

	ownerResp := &OwnerResp{
		Address: result.Owner.Address,
		IDCs:    idcs,
	}
	response := &QueryOwnerResp{
		OwnerResp: ownerResp,
		Pagination: &PageResponse{
			NextKey: result.Pagination.NextKey,
			Total:   result.Pagination.Total,
		},
	}
	return response, nil
}

func (cli *Client) QueryCollection(classId string, pagination PaginationRequest) (*QueryCollectionResp, error) {

	if len(classId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidDenom, "denom is required")
	}

	page := &query.PageRequest{
		CountTotal: false,
		Offset:     pagination.Offset,
	}

	if pagination.Limit > 0 {
		page.Limit = pagination.Limit
	} else {
		page.Limit = 100
	}

	if len(pagination.NextKey) > 0 {
		page.Key = []byte(pagination.NextKey)
	}

	req := &QueryCollectionRequest{
		DenomId:    classId,
		Pagination: page,
	}

	result, err := cli.queryCli.Collection(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to query collection: %s", err)
	}

	queryClassResp := &QueryClassResp{
		ID:               result.Collection.Denom.Id,
		Name:             result.Collection.Denom.Name,
		Schema:           result.Collection.Denom.Schema,
		Creator:          result.Collection.Denom.Creator,
		Symbol:           result.Collection.Denom.Symbol,
		MintRestricted:   result.Collection.Denom.MintRestricted,
		UpdateRestricted: result.Collection.Denom.UpdateRestricted,
		Data:             result.Collection.Denom.Data,
		Uri:              result.Collection.Denom.Uri,
		UriHash:          result.Collection.Denom.UriHash,
		Description:      result.Collection.Denom.Description,
	}
	var nfts []QueryNFTResp
	for _, value := range result.Collection.NFTs {
		nft := QueryNFTResp{
			ID:      value.Id,
			Name:    value.Name,
			Owner:   value.Owner,
			Data:    value.Data,
			URI:     value.URI,
			URIHash: value.UriHash,
		}
		nfts = append(nfts, nft)
	}

	response := &QueryCollectionResp{
		Class: queryClassResp,
		NFTs:  nfts,
		Pagination: &PageResponse{
			NextKey: result.Pagination.NextKey,
			Total:   result.Pagination.Total,
		},
	}

	return response, nil
}

func (cli *Client) QueryClass(classId string) (*QueryClassResp, error) {
	if len(classId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidDenom, "class id is required")
	}

	result, err := cli.queryCli.Denom(
		context.Background(),
		&QueryDenomRequest{
			DenomId: classId,
		})
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to query class: %s", err)
	}

	response := &QueryClassResp{
		ID:               result.Denom.Id,
		Name:             result.Denom.Name,
		Schema:           result.Denom.Schema,
		Creator:          result.Denom.Creator,
		Symbol:           result.Denom.Symbol,
		MintRestricted:   result.Denom.MintRestricted,
		UpdateRestricted: result.Denom.UpdateRestricted,
		Data:             result.Denom.Data,
		Uri:              result.Denom.Uri,
		UriHash:          result.Denom.UriHash,
		Description:      result.Denom.Description,
	}

	return response, nil
}

func (cli *Client) QueryClasses(pagination PaginationRequest) (*QueryClassesResp, error) {
	page := &query.PageRequest{
		CountTotal: false,
		Offset:     pagination.Offset,
	}

	if pagination.Limit > 0 {
		page.Limit = pagination.Limit
	} else {
		page.Limit = 100
	}

	if len(pagination.NextKey) > 0 {
		page.Key = []byte(pagination.NextKey)
	}

	result, err := cli.queryCli.Denoms(
		context.Background(),
		&QueryDenomsRequest{
			Pagination: page,
		})
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to query classes: %s", err)
	}
	var classes []QueryClassResp
	for _, value := range result.Denoms {
		class := QueryClassResp{
			ID:               value.Id,
			Name:             value.Name,
			Schema:           value.Schema,
			Creator:          value.Creator,
			Symbol:           value.Symbol,
			MintRestricted:   value.MintRestricted,
			UpdateRestricted: value.UpdateRestricted,
			Data:             value.Data,
			Uri:              value.Uri,
			UriHash:          value.UriHash,
			Description:      value.Description,
		}
		classes = append(classes, class)
	}

	response := &QueryClassesResp{
		Classes: classes,
		Pagination: &PageResponse{
			NextKey: result.Pagination.NextKey,
			Total:   result.Pagination.Total,
		},
	}

	return response, nil
}

func (cli *Client) QueryNFT(classId, tokenID string) (*QueryNFTResp, error) {
	if len(classId) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidDenom, "class id is required")
	}

	if len(tokenID) == 0 {
		return nil, sdkerrors.Wrapf(ErrInvalidTokenID, "token id is required")
	}
	req := &QueryNFTRequest{
		DenomId: classId,
		TokenId: tokenID,
	}
	result, err := cli.queryCli.NFT(
		context.Background(),
		req,
	)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "failed to query nft: %s", err)
	}

	response := &QueryNFTResp{
		ID:      result.NFT.Id,
		Name:    result.NFT.Name,
		Owner:   result.NFT.Owner,
		Data:    result.NFT.Data,
		URI:     result.NFT.URI,
		URIHash: result.NFT.UriHash,
	}

	return response, nil
}
