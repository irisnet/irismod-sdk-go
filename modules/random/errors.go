package random

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + ModuleName

var (
	ErrInvalidReqID         = sdkerrors.Register(Codespace, 1, " length error")
	ErrGenConn              = sdkerrors.Register(Codespace, 2, "generate conn error")
	ErrAtoi                 = sdkerrors.Register(Codespace, 3, "strconv.Atoi error")
	ErrAccAddressFromBech32 = sdkerrors.Register(Codespace, 4, "AccAddressFromBech32 error ")
	ErrQueryRandom          = sdkerrors.Register(Codespace, 5, "QueryRandom error ")
	ErrEventsGetValue       = sdkerrors.Register(Codespace, 6, "EventsGetValue error ")
	ErrQueryRequestQueue    = sdkerrors.Register(Codespace, 7, "QueryRequestQueue error ")
)
