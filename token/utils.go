package token

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"
)

var (
	nativeToken Token
	Initialized bool
)

// GetNativeToken returns the system's default native token
func GetNativeToken() Token {
	if !Initialized {
		nativeToken = Token{
			Symbol:        sdk.DefaultBondDenom,
			Name:          "Network staking token",
			Scale:         0,
			MinUnit:       sdk.DefaultBondDenom,
			InitialSupply: 2000000000,
			MaxSupply:     10000000000,
			Mintable:      true,
			Owner:         sdk.AccAddress(crypto.AddressHash([]byte(ModuleName))).String(),
		}
		Initialized = true
	}
	return nativeToken
}
