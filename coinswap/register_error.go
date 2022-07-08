package coinswap

import (
	sdkerrors "github.com/irisnet/core-sdk-go/types"
)

// coinswap module sentinel errors
var (
	ErrReservePoolNotExists    = sdkerrors.RegisterErr(ModuleName, 2, "reserve pool not exists")
	ErrEqualDenom              = sdkerrors.RegisterErr(ModuleName, 3, "input and output denomination are equal")
	ErrNotContainStandardDenom = sdkerrors.RegisterErr(ModuleName, 4, "must have one standard denom")
	ErrMustStandardDenom       = sdkerrors.RegisterErr(ModuleName, 5, "must be standard denom")
	ErrInvalidDenom            = sdkerrors.RegisterErr(ModuleName, 6, "invalid denom")
	ErrInvalidDeadline         = sdkerrors.RegisterErr(ModuleName, 7, "invalid deadline")
	ErrConstraintNotMet        = sdkerrors.RegisterErr(ModuleName, 8, "constraint not met")
	ErrInsufficientFunds       = sdkerrors.RegisterErr(ModuleName, 9, "insufficient funds")
)
