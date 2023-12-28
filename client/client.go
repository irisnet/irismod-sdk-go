package client

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/irisnet/core-sdk-go/client"
	sdk "github.com/irisnet/core-sdk-go/types"
	"github.com/irisnet/irismod-sdk-go/mt"
	"github.com/irisnet/irismod-sdk-go/nft"
	"github.com/tendermint/tendermint/libs/log"
)

type Client struct {
	moduleManager  map[string]sdk.Module
	encodingConfig sdk.EncodingConfig
	sdk.BaseClient

	NFTClient nft.IClient
	MTClient  mt.IClient
}

func NewClient(cfg sdk.ClientConfig) (Client, error) {
	encodingConfig := sdk.MakeEncodingConfig()

	// create a instance of baseClient
	baseClient := client.NewBaseClient(cfg, encodingConfig)

	// create irismod client
	nftClient, err := nft.NewClient(baseClient)
	if err != nil {
		return Client{}, err
	}

	mtClient, err := mt.NewClient(baseClient)
	if err != nil {
		return Client{}, err
	}

	cli := Client{
		BaseClient:     baseClient,
		moduleManager:  make(map[string]sdk.Module),
		encodingConfig: encodingConfig,
		NFTClient:      nftClient,
		MTClient:       mtClient,
	}
	cli.RegisterModule(nftClient, mtClient)
	return cli, nil
}

func (client *Client) SetLogger(logger log.Logger) {
	client.BaseClient.SetLogger(logger)
}

func (client *Client) Codec() *codec.LegacyAmino {
	return client.encodingConfig.Amino
}

func (client *Client) AppCodec() codec.Codec {
	return client.encodingConfig.Marshaler
}

func (client *Client) EncodingConfig() sdk.EncodingConfig {
	return client.encodingConfig
}

func (client *Client) Manager() sdk.BaseClient {
	return client.BaseClient
}

func (client *Client) RegisterModule(ms ...sdk.Module) {
	for _, m := range ms {
		m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
	}
}

func (client *Client) Module(name string) sdk.Module {
	return client.moduleManager[name]
}
