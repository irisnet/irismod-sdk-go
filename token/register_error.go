package token

import (
	sdkerrors "github.com/irisnet/core-sdk-go/types"
)

// token module sentinel errors
var (
	ErrInvalidName          = sdkerrors.RegisterErr(ModuleName, 2, "invalid token name")
	ErrInvalidMinUnit       = sdkerrors.RegisterErr(ModuleName, 3, "invalid token min unit")
	ErrInvalidSymbol        = sdkerrors.RegisterErr(ModuleName, 4, "invalid standard denom")
	ErrInvalidInitSupply    = sdkerrors.RegisterErr(ModuleName, 5, "invalid token initial supply")
	ErrInvalidMaxSupply     = sdkerrors.RegisterErr(ModuleName, 6, "invalid token maximum supply")
	ErrInvalidScale         = sdkerrors.RegisterErr(ModuleName, 7, "invalid token scale")
	ErrSymbolAlreadyExists  = sdkerrors.RegisterErr(ModuleName, 8, "symbol already exists")
	ErrMinUnitAlreadyExists = sdkerrors.RegisterErr(ModuleName, 9, "min unit already exists")
	ErrTokenNotExists       = sdkerrors.RegisterErr(ModuleName, 10, "token does not exist")
	ErrInvalidToAddress     = sdkerrors.RegisterErr(ModuleName, 11, "the new owner must not be same as the original owner")
	ErrInvalidOwner         = sdkerrors.RegisterErr(ModuleName, 12, "invalid token owner")
	ErrNotMintable          = sdkerrors.RegisterErr(ModuleName, 13, "token is not mintable")
	ErrNotFoundTokenAmt     = sdkerrors.RegisterErr(ModuleName, 14, "burned token amount not found")
	ErrInvalidAmount        = sdkerrors.RegisterErr(ModuleName, 15, "invalid amount")
	ErrInvalidBaseFee       = sdkerrors.RegisterErr(ModuleName, 16, "invalid base fee")
)
