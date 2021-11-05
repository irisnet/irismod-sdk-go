package modules

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace

var (
	ErrQueryToken       = sdkerrors.Register(Codespace, 1, "QueryToken error")
	ErrConvertToMinCoin = sdkerrors.Register(Codespace, 2, "query address error")
	ErrToMintCoin       = sdkerrors.Register(Codespace, 3, "ConvertToMinCoin error")
	ErrUnpackAny        = sdkerrors.Register(Codespace, 4, " UnpackAny error")
	ErrGenConn          = sdkerrors.Register(Codespace, 5, "generate conn error")
)
