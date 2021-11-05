package record

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + ModuleName

var (
	ErrQueryAddress       = sdkerrors.Register(Codespace, 1, "query address error")
	ErrDecodeString       = sdkerrors.Register(Codespace, 2, "decode string error")
	ErrValidateAccAddress = sdkerrors.Register(Codespace, 3, "ValidateAccAddress error ")
	ErrQueryStore         = sdkerrors.Register(Codespace, 4, "query Random value error ")
	ErrEventsGetValue     = sdkerrors.Register(Codespace, 5, "query Random value error ")
	ErrCodecUnmarshal     = sdkerrors.Register(Codespace, 6, "query Random value error ")
)
