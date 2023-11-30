package mt

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrInvalidClassId = sdkerrors.Register(ModuleName, 19, "invalid classId")
	ErrInvalidTokenID = sdkerrors.Register(ModuleName, 20, "invalid tokenId")
	ErrInvalidHeight  = sdkerrors.Register(ModuleName, 21, "invalid height")
)
