package random

import (
	sdkerrors "github.com/irisnet/core-sdk-go/types"
)

// random module sentinel errors
var (
	ErrInvalidReqID            = sdkerrors.RegisterErr(ModuleName, 2, "invalid request id")
	ErrInvalidHeight           = sdkerrors.RegisterErr(ModuleName, 3, "invalid height, must be greater than 0")
	ErrInvalidServiceBindings  = sdkerrors.RegisterErr(ModuleName, 4, "no service bindings available")
	ErrInvalidRequestContextID = sdkerrors.RegisterErr(ModuleName, 5, "invalid request context id")
	ErrInvalidServiceFeeCap    = sdkerrors.RegisterErr(ModuleName, 6, "invalid service fee cap")
)
