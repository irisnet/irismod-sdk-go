# IRISMOD SDK
Golang SDK for IRISnet Modules

## Requirement
Go version above 1.16.4

## Use Go Mod

```
replace (
    github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
    github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
```

## Customize the Client type

```go
    type Client struct {
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
```
You can define a client according to your own needs


## New Client

```go
    func NewClient(cfg types.ClientConfig) Client {
        encodingConfig := makeEncodingConfig()
    
        // create a instance of baseClient
        baseClient := client.NewBaseClient(cfg, encodingConfig, nil)
    
        keysClient := client.NewKeysClient(cfg, baseClient)
    
        bankClient := bank.NewClient(baseClient, encodingConfig.Marshaler)
        tokenClient := token.NewClient(baseClient, encodingConfig.Marshaler)
        stakingClient := staking.NewClient(baseClient, encodingConfig.Marshaler)
        govClient := gov.NewClient(baseClient, encodingConfig.Marshaler)
        serviceClient := service.NewClient(baseClient, encodingConfig.Marshaler)
        recordClient := record.NewClient(baseClient, encodingConfig.Marshaler)
        nftClient := nft.NewClient(baseClient, encodingConfig.Marshaler)
        randomClient := random.NewClient(baseClient, encodingConfig.Marshaler)
        oracleClient := oracle.NewClient(baseClient, encodingConfig.Marshaler)
        htlcClient := htlc.NewClient(baseClient, encodingConfig.Marshaler)
        swapClient := coinswap.NewClient(baseClient, encodingConfig.Marshaler, bankClient.TotalSupply)
    
        client := &Client{
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

```

## Module Register

```go
    func (client Client) RegisterModule(ms ...types.Module) {
        for _, m := range ms {
            m.RegisterInterfaceTypes(client.encodingConfig.InterfaceRegistry)
        }
    }

```


## Init Client
The initialization SDK code is as follows:

```go
    options := []types.Option{
            types.KeyDAOOption(store.NewMemory(nil)),
            types.TimeoutOption(10),
        }
        cfg, err := types.NewClientConfig(nodeURI, grpcAddr, chainID, options...)
        if err != nil {
            panic(err)
        }
    
        s.Client = NewClient(cfg)
```


