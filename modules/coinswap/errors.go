package coinswap

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + ModuleName

var (
	ErrQueryAddress         = sdkerrors.Register(Codespace, 1, "query address error")
	ErrConvertInt           = sdkerrors.Register(Codespace, 2, "convert to sdk.Int error")
	ErrGenConn              = sdkerrors.Register(Codespace, 3, "generate conn error")
	ErrLiquidityPool        = sdkerrors.Register(Codespace, 4, "query liquidity pool error")
	ErrDenom                = sdkerrors.Register(Codespace, 5, "denom error")
	ErrPrefix               = sdkerrors.Register(Codespace, 6, "prefix error")
	ErrTokenCount           = sdkerrors.Register(Codespace, 7, "invalid token count")
	ErrDeadline             = sdkerrors.Register(Codespace, 8, "invalid deadline")
	ErrAccAddressFromBech32 = sdkerrors.Register(Codespace, 9, "AccAddressFromBech32 error ")
)
