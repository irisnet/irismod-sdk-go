package mt

import (
	context "context"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/types"
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

func (cli *Client) CreateClass(request IssueClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgIssueDenom{
		Name:   request.Name,
		Sender: sender.String(),
		Data:   request.Data,
	}
	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

// MintMT create new MT
func (cli *Client) MintMT(request MintMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	var recipient = sender.String()
	if len(request.Recipient) > 0 {
		if _, err := types.ValAddressFromBech32(request.Recipient); err != nil {
			return sdk.ResultTx{}, sdk.Wrap(err)
		}
		recipient = request.Recipient
	}

	msg := &MsgMintMT{
		DenomId:   request.ClassID,
		Amount:    request.Amount,
		Data:      request.Data,
		Sender:    sender.String(),
		Recipient: recipient,
	}
	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

// AddMT issuing additional MT
func (cli *Client) AddMT(request AddMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	var recipient = sender.String()
	if len(request.Recipient) > 0 {
		if _, err := types.ValAddressFromBech32(request.Recipient); err != nil {
			return sdk.ResultTx{}, sdk.Wrap(err)
		}
		recipient = request.Recipient
	}

	msg := &MsgMintMT{
		Id:        request.ID,
		DenomId:   request.ClassID,
		Amount:    request.Amount,
		Sender:    sender.String(),
		Recipient: recipient,
	}
	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) EditMT(request EditMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgEditMT{
		Id:      request.ID,
		DenomId: request.ClassID,
		Data:    request.Data,
		Sender:  sender.String(),
	}
	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) TransferMT(request TransferMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	if _, err := types.ValAddressFromBech32(request.Recipient); err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgTransferMT{
		Id:        request.ID,
		DenomId:   request.ClassID,
		Amount:    request.Amount,
		Sender:    sender.String(),
		Recipient: request.Recipient,
	}
	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) BurnMT(request BurnMTRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgBurnMT{
		Id:      request.ID,
		DenomId: request.ClassID,
		Amount:  request.Amount,
		Sender:  sender.String(),
	}
	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) TransferClass(request TransferClassRequest, baseTx sdk.BaseTx) (sdk.ResultTx, sdk.Error) {
	sender, err := cli.QueryAddress(baseTx.From, baseTx.Password)
	if err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	if _, err := types.ValAddressFromBech32(request.Recipient); err != nil {
		return sdk.ResultTx{}, sdk.Wrap(err)
	}

	msg := &MsgTransferDenom{
		Id:        request.ID,
		Sender:    sender.String(),
		Recipient: request.Recipient,
	}
	return cli.BuildAndSend([]types.Msg{msg}, baseTx)
}

func (cli *Client) QuerySupply(denomID, creator string) (uint64, sdk.Error) {
	if len(denomID) == 0 {
		return 0, sdk.Wrapf("denomID is required")
	}
	if _, err := types.ValAddressFromBech32(creator); err != nil {
		return 0, sdk.Wrap(err)
	}

	res, err := cli.queryCli.Supply(
		context.Background(),
		&QuerySupplyRequest{
			Owner:   creator,
			DenomId: denomID,
		},
	)
	if err != nil {
		return 0, sdk.Wrap(err)
	}

	return res.Amount, nil
}

func (cli *Client) QueryMTSupply(denomID, mtID string) (uint64, sdk.Error) {
	if len(denomID) == 0 {
		return 0, sdk.Wrapf("denomID is required")
	}
	if len(mtID) == 0 {
		return 0, sdk.Wrapf("mtID is required")
	}

	res, err := cli.queryCli.MTSupply(
		context.Background(),
		&QueryMTSupplyRequest{
			DenomId: denomID,
			MtId:    mtID,
		},
	)
	if err != nil {
		return 0, sdk.Wrap(err)
	}

	return res.Amount, nil
}

func (cli *Client) QueryClass(denomID string) (QueryClassResp, sdk.Error) {
	if len(denomID) == 0 {
		return QueryClassResp{}, sdk.Wrapf("denomID is required")
	}

	res, err := cli.queryCli.Denom(
		context.Background(),
		&QueryDenomRequest{DenomId: denomID},
	)
	if err != nil {
		return QueryClassResp{}, sdk.Wrap(err)
	}
	return QueryClassResp{
		ID:    res.Denom.Id,
		Name:  res.Denom.Name,
		Data:  res.Denom.Data,
		Owner: res.Denom.Owner,
	}, nil

}

func (cli *Client) QueryClasses(pagination *query.PageRequest) ([]QueryClassResp, sdk.Error) {

	res, err := cli.queryCli.Denoms(
		context.Background(),
		&QueryDenomsRequest{
			Pagination: pagination,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	var denoms []QueryClassResp
	for _, denom := range res.Denoms {
		denoms = append(denoms, QueryClassResp{
			ID:    denom.Id,
			Name:  denom.Name,
			Data:  denom.Data,
			Owner: denom.Owner,
		})
	}
	return denoms, nil
}

func (cli *Client) QueryMT(denomID, mtID string) (QueryMTResp, sdk.Error) {
	if len(denomID) == 0 {
		return QueryMTResp{}, sdk.Wrapf("denomID is required")
	}

	if len(mtID) == 0 {
		return QueryMTResp{}, sdk.Wrapf("mtID is required")
	}

	res, err := cli.queryCli.MT(
		context.Background(),
		&QueryMTRequest{
			DenomId: denomID,
			MtId:    mtID,
		},
	)
	if err != nil {
		return QueryMTResp{}, sdk.Wrap(err)
	}

	return QueryMTResp{
		ID:     res.Mt.Id,
		Supply: res.Mt.Supply,
		Data:   res.Mt.Data,
	}, nil
}

func (cli *Client) QueryMTs(denomID string, pagination *query.PageRequest) ([]QueryMTResp, sdk.Error) {
	if len(denomID) == 0 {
		return nil, sdk.Wrapf("denomID is required")
	}

	res, err := cli.queryCli.MTs(
		context.Background(),
		&QueryMTsRequest{
			DenomId:    denomID,
			Pagination: pagination,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	var mts []QueryMTResp
	for _, mt := range res.Mts {
		mts = append(mts, QueryMTResp{
			ID:     mt.Id,
			Supply: mt.Supply,
			Data:   mt.Data,
		})
	}

	return mts, nil
}

func (cli *Client) QueryBalances(denomID, owner string) ([]QueryBalanceResp, sdk.Error) {
	if len(denomID) == 0 {
		return nil, sdk.Wrapf("denomID is required")
	}

	if _, err := types.ValAddressFromBech32(owner); err != nil {
		return nil, sdk.Wrap(err)
	}

	res, err := cli.queryCli.Balances(
		context.Background(),
		&QueryBalancesRequest{
			DenomId: denomID,
			Owner:   owner,
		},
	)
	if err != nil {
		return nil, sdk.Wrap(err)
	}
	var balances []QueryBalanceResp
	for _, balance := range res.Balance {
		balances = append(balances, QueryBalanceResp{
			MtId:   balance.MtId,
			Amount: balance.Amount,
		})
	}
	return balances, nil
}
