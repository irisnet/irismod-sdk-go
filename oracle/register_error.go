package oracle

import (
	sdkerrors "github.com/irisnet/core-sdk-go/types"
)

// oracle module sentinel errors
var (
	ErrUnknownFeedName      = sdkerrors.RegisterErr(ModuleName, 2, "unknown feed")
	ErrInvalidFeedName      = sdkerrors.RegisterErr(ModuleName, 3, "invalid feed name")
	ErrExistedFeedName      = sdkerrors.RegisterErr(ModuleName, 4, "feed already exists")
	ErrUnauthorized         = sdkerrors.RegisterErr(ModuleName, 5, "unauthorized owner")
	ErrInvalidServiceName   = sdkerrors.RegisterErr(ModuleName, 6, "invalid service name")
	ErrInvalidDescription   = sdkerrors.RegisterErr(ModuleName, 7, "invalid description")
	ErrNotRegisterFunc      = sdkerrors.RegisterErr(ModuleName, 8, "method don't register")
	ErrInvalidFeedState     = sdkerrors.RegisterErr(ModuleName, 9, "invalid state feed")
	ErrInvalidServiceFeeCap = sdkerrors.RegisterErr(ModuleName, 10, "service fee cap is invalid")
)
