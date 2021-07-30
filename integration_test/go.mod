module github.com/irisnet/irismod-sdk-go/integration-test

go 1.16

require (
	github.com/irisnet/core-sdk-go v0.0.0-20210719031639-9c6ece68d908
	github.com/irisnet/irismod-sdk-go/coinswap v0.0.0-20210729071729-a138c14c92e5
	github.com/irisnet/irismod-sdk-go/farm v0.0.0-20210729071729-a138c14c92e5
	github.com/irisnet/irismod-sdk-go/htlc v0.0.0-20210729071729-a138c14c92e5
	github.com/irisnet/irismod-sdk-go/nft v0.0.0-20210729071729-a138c14c92e5
	github.com/irisnet/irismod-sdk-go/oracle v0.0.0-20210729071729-a138c14c92e5
	github.com/irisnet/irismod-sdk-go/random v0.0.0-20210729071729-a138c14c92e5
	github.com/irisnet/irismod-sdk-go/record v0.0.0-20210729071729-a138c14c92e5
	github.com/irisnet/irismod-sdk-go/service v0.0.0-20210729071729-a138c14c92e5
	github.com/irisnet/irismod-sdk-go/token v0.0.0-20210729071729-a138c14c92e5
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
