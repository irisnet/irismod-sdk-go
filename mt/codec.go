package mt

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

// RegisterLegacyAminoCodec concrete types on codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgIssueDenom{}, "irismod/mt/MsgIssueDenom", nil)
	cdc.RegisterConcrete(&MsgTransferMT{}, "irismod/mt/MsgTransferMT", nil)
	cdc.RegisterConcrete(&MsgEditMT{}, "irismod/mt/MsgEditMT", nil)
	cdc.RegisterConcrete(&MsgMintMT{}, "irismod/mt/MsgMintMT", nil)
	cdc.RegisterConcrete(&MsgBurnMT{}, "irismod/mt/MsgBurnMT", nil)
	cdc.RegisterConcrete(&MsgTransferDenom{}, "irismod/mt/MsgTransferDenom", nil)

}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgIssueDenom{},
		&MsgTransferMT{},
		&MsgEditMT{},
		&MsgMintMT{},
		&MsgBurnMT{},
		&MsgTransferDenom{},
	)

}
