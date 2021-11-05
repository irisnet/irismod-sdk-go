package keys

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + "keys"

var (
	ErrInsert  = sdkerrors.Register(Codespace, 1, "key-manager insert error")
	ErrRecover = sdkerrors.Register(Codespace, 2, "key-manager recover error")
	ErrImport  = sdkerrors.Register(Codespace, 3, "key-manager import error")
	ErrExport  = sdkerrors.Register(Codespace, 4, "key-manager export error")
	ErrDelete  = sdkerrors.Register(Codespace, 5, "key-manager delete error")
	ErrShow    = sdkerrors.Register(Codespace, 6, "key-manager show error")
)
