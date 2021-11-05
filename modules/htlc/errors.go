package htlc

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + ModuleName

var (
	ErrInvalidRequest       = sdkerrors.Register(Codespace, 1, "invalid request")
	ErrQueryAddress         = sdkerrors.Register(Codespace, 2, "query address error")
	ErrToMintCoin           = sdkerrors.Register(Codespace, 3, " to mint coin error")
	ErrInvalidLength        = sdkerrors.Register(Codespace, 4, " length error")
	ErrGenConn              = sdkerrors.Register(Codespace, 5, "generate conn error")
	ErrQueryHTLC            = sdkerrors.Register(Codespace, 6, "query htlc error")
	ErrQueryParams          = sdkerrors.Register(Codespace, 7, "query params error")
	ErrDecodeString         = sdkerrors.Register(Codespace, 8, "decode string error")
	ErrTimeLock             = sdkerrors.Register(Codespace, 9, "invalid time lock")
	ErrTokenCount           = sdkerrors.Register(Codespace, 10, "invalid token count")
	ErrAccAddressFromBech32 = sdkerrors.Register(Codespace, 11, "AccAddressFromBech32 error ")
)
