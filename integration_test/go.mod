module github.com/irisnet/irismod-sdk-go/integration-test

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.1 // indirect
	github.com/irisnet/core-sdk-go v0.0.0-20210719031639-9c6ece68d908
	github.com/irisnet/irismod-sdk-go/coinswap v0.0.0-20210720100738-ddc220cd5bd1
	github.com/irisnet/irismod-sdk-go/gov v0.0.0-20210720100738-ddc220cd5bd1
	github.com/irisnet/irismod-sdk-go/htlc v0.0.0-20210720100738-ddc220cd5bd1
	github.com/irisnet/irismod-sdk-go/mt v0.0.0-20220304092109-d22616e3da29
	github.com/irisnet/irismod-sdk-go/nft v0.0.0-20210726064324-b35f8f5259eb
	github.com/irisnet/irismod-sdk-go/oracle v0.0.0-20210720100738-ddc220cd5bd1
	github.com/irisnet/irismod-sdk-go/random v0.0.0-20210720100738-ddc220cd5bd1
	github.com/irisnet/irismod-sdk-go/record v0.0.0-20210720100738-ddc220cd5bd1
	github.com/irisnet/irismod-sdk-go/service v0.0.0-20210720100133-4b2a0a8cc4f1
	github.com/irisnet/irismod-sdk-go/staking v0.0.0-20210720100738-ddc220cd5bd1
	github.com/irisnet/irismod-sdk-go/token v0.0.0-20210720100738-ddc220cd5bd1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.14
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/prometheus/common => github.com/prometheus/common v0.26.0
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
