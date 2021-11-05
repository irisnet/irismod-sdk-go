package nft

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + ModuleName

var (
	ErrQueryAddress       = sdkerrors.Register(Codespace, 1, "query address error")
	ErrInvalidDenom       = sdkerrors.Register(Codespace, 2, " to mint coin error")
	ErrInvalidTokenID     = sdkerrors.Register(Codespace, 3, " to mint coin error")
	ErrGenConn            = sdkerrors.Register(Codespace, 4, "generate conn error")
	ErrInvalidAddress     = sdkerrors.Register(Codespace, 5, "invalid address")
	ErrValidateAccAddress = sdkerrors.Register(Codespace, 6, "ValidateAccAddress error ")
	ErrQuerySupply        = sdkerrors.Register(Codespace, 7, "query supply error ")
	ErrQueryOwner         = sdkerrors.Register(Codespace, 8, "query owner error ")
	ErrQueryCollection    = sdkerrors.Register(Codespace, 9, "query collection error ")
	ErrQueryDenoms        = sdkerrors.Register(Codespace, 10, "query denoms error ")
	ErrQueryNFT           = sdkerrors.Register(Codespace, 11, "query nft error ")
)
