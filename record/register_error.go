package record

import (
	sdkerrors "github.com/irisnet/core-sdk-go/types"
)

// record module sentinel errors
var (
	ErrUnknownRecord = sdkerrors.RegisterErr(ModuleName, 2, "unknown record")
)
