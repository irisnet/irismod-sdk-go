module github.com/irisnet/irismod-sdk-go/nft

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.45.1
	github.com/gogo/protobuf v1.3.3
	github.com/irisnet/core-sdk-go v0.0.0-20220302175731-8770d7dce833
	google.golang.org/genproto v0.0.0-20210828152312-66f60bf46e71
	google.golang.org/grpc v1.42.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
