module github.com/irisnet/irismod-sdk-go/random

go 1.16

require (
	github.com/gogo/protobuf v1.3.3
	github.com/irisnet/core-sdk-go v0.0.0-20220708022742-798813c6ab8e
	github.com/tendermint/tendermint v0.34.19 // indirect
	google.golang.org/genproto v0.0.0-20211116182654-e63d96a377c4
	google.golang.org/grpc v1.41.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
