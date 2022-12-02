module github.com/irisnet/irismod-sdk-go/htlc

go 1.16

require (
	github.com/gogo/protobuf v1.3.3
	github.com/irisnet/core-sdk-go v0.0.0-20220720085949-4d825adb8054
	google.golang.org/genproto v0.0.0-20211116182654-e63d96a377c4
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
