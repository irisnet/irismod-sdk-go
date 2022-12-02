module github.com/irisnet/irismod-sdk-go/integration-test

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.1 // indirect
	github.com/irisnet/core-sdk-go v0.0.0-20220720085949-4d825adb8054
	github.com/irisnet/irismod-sdk-go/coinswap v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/gov v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/htlc v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/mt v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/nft v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/oracle v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/random v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/record v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/service v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/staking v0.0.0-20221202022845-2fad43fde95d
	github.com/irisnet/irismod-sdk-go/token v0.0.0-20221202022845-2fad43fde95d
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.19
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/irisnet/irismod-sdk-go/coinswap => ../coinswap
	github.com/irisnet/irismod-sdk-go/gov => ../gov
	github.com/irisnet/irismod-sdk-go/htlc => ../htlc
	github.com/irisnet/irismod-sdk-go/mt => ../mt
	github.com/irisnet/irismod-sdk-go/nft => ../nft
	github.com/irisnet/irismod-sdk-go/oracle => ../oracle
	github.com/irisnet/irismod-sdk-go/random => ../random
	github.com/irisnet/irismod-sdk-go/record => ../record
	github.com/irisnet/irismod-sdk-go/service => ../service
	github.com/irisnet/irismod-sdk-go/staking => ../staking
	github.com/irisnet/irismod-sdk-go/token => ../token
	github.com/prometheus/common => github.com/prometheus/common v0.26.0
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
