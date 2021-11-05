package sdk

import (
	"fmt"

	"github.com/irisnet/core-sdk-go/client"

	"github.com/irisnet/core-sdk-go/modules/gov"
	"github.com/irisnet/core-sdk-go/modules/staking"
	"github.com/irisnet/irishub-sdk-go/modules/coinswap"
	"github.com/irisnet/irishub-sdk-go/modules/htlc"
	"github.com/irisnet/irishub-sdk-go/modules/nft"
	"github.com/irisnet/irishub-sdk-go/modules/oracle"
	"github.com/irisnet/irishub-sdk-go/modules/random"
	"github.com/irisnet/irishub-sdk-go/modules/record"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/irisnet/core-sdk-go/codec"
	cdctypes "github.com/irisnet/core-sdk-go/codec/types"
	cryptocodec "github.com/irisnet/core-sdk-go/crypto/codec"
	"github.com/irisnet/core-sdk-go/modules/bank"
	"github.com/irisnet/core-sdk-go/types"
	txtypes "github.com/irisnet/core-sdk-go/types/tx"
	"github.com/irisnet/irishub-sdk-go/modules/keys"
	"github.com/irisnet/irishub-sdk-go/modules/service"
	"github.com/irisnet/irishub-sdk-go/modules/token"
)

type IRISHUBClient struct {
	logger         log.Logger
	moduleManager  map[string]types.Module
	encodingConfig types.EncodingConfig

	types.BaseClient
	Key     keys.Client
	Bank    bank.Client
	Token   token.Client
	Staking staking.Client
	Gov     gov.Client
	Service service.Client
	Record  record.Client
	Random  random.Client
	NFT     nft.Client
	Oracle  oracle.Client
	HTLC    htlc.Client
	Swap    coinswap.Client
}

func NewIRISHUBClient(cfg types.ClientConfig) IRISHUBClient {
	encodingConfig := makeEncodingConfig()
	// create a instance of baseClient
	baseClient := client.NewBaseClient(cfg, encodingConfig, nil)
	keysClient := keys.NewClient(baseClient)

	bankClient := bank.NewClient(baseClient, encodingConfig.Codec)
	tokenClient := token.NewClient(baseClient, encodingConfig.Codec)
	stakingClient := staking.NewClient(baseClient, encodingConfig.Codec)
	govClient := gov.NewClient(baseClient, encodingConfig.Codec)

	serviceClient := service.NewClient(baseClient, encodingConfig.Codec)
	recordClient := record.NewClient(baseClient, encodingConfig.Codec)
	nftClient := nft.NewClient(baseClient, encodingConfig.Codec)
	randomClient := random.NewClient(baseClient, encodingConfig.Codec)
	oracleClient := oracle.NewClient(baseClient, encodingConfig.Codec)
	htlcClient := htlc.NewClient(baseClient, encodingConfig.Codec)
	swapClient := coinswap.NewClient(baseClient, encodingConfig.Codec, bankClient.TotalSupply)

	client := &IRISHUBClient{
		logger:         baseClient.Logger(),
		BaseClient:     baseClient,
		moduleManager:  make(map[string]types.Module),
		encodingConfig: encodingConfig,
		Key:            keysClient,
		Bank:           bankClient,
		Token:          tokenClient,
		Staking:        stakingClient,
		Gov:            govClient,
		Service:        serviceClient,
		Record:         recordClient,
		Random:         randomClient,
		NFT:            nftClient,
		Oracle:         oracleClient,
		HTLC:           htlcClient,
		Swap:           swapClient,
	}

	client.RegisterModule(
		bankClient,
		tokenClient,
		stakingClient,
		govClient,
		serviceClient,
		recordClient,
		nftClient,
		randomClient,
		oracleClient,
		htlcClient,
		swapClient,
	)
	return *client
}

func (client *IRISHUBClient) SetLogger(logger log.Logger) {
	client.BaseClient.SetLogger(logger)
}

func (client *IRISHUBClient) Codec() codec.Codec {
	return client.encodingConfig.Codec
}

func (client *IRISHUBClient) AppCodec() codec.Codec {
	return client.encodingConfig.Codec
}

func (client *IRISHUBClient) EncodingConfig() types.EncodingConfig {
	return client.encodingConfig
}

func (client *IRISHUBClient) Manager() types.BaseClient {
	return client.BaseClient
}

func (client *IRISHUBClient) RegisterModule(ms ...types.Module) {
	for _, m := range ms {
		_, ok := client.moduleManager[m.Name()]
		if ok {
			panic(fmt.Sprintf("%s has register", m.Name()))
		}

		// m.RegisterCodec(client.encodingConfig.Amino)
		m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
		client.moduleManager[m.Name()] = m
	}
}

func (client *IRISHUBClient) Module(name string) types.Module {
	return client.moduleManager[name]
}

func makeEncodingConfig() types.EncodingConfig {
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txCfg := txtypes.NewTxConfig(marshaler, txtypes.DefaultSignModes)

	encodingConfig := types.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             marshaler,
		TxConfig:          txCfg,
	}
	RegisterInterfaces(encodingConfig.InterfaceRegistry)
	return encodingConfig
}

// RegisterLegacyAminoCodec registers the sdk message type.
//func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
//	cdc.RegisterInterface((*types.Msg)(nil), nil)
//	cdc.RegisterInterface((*types.Tx)(nil), nil)
//	cryptocodec.RegisterCrypto(cdc)
//}

// RegisterInterfaces registers the sdk message type.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterInterface("cosmos.v1beta1.Msg", (*types.Msg)(nil))
	txtypes.RegisterInterfaces(registry)
	cryptocodec.RegisterInterfaces(registry)
}
