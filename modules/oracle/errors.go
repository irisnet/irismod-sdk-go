package oracle

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + ModuleName

var (
	ErrQueryAddress         = sdkerrors.Register(Codespace, 1, "query address error")
	ErrToMintCoin           = sdkerrors.Register(Codespace, 2, " to mint coin error")
	ErrInvalidServiceName   = sdkerrors.Register(Codespace, 3, "invalid service name")
	ErrValueJsonPath        = sdkerrors.Register(Codespace, 4, "invalid valueJsonPath")
	ErrServiceFeeCap        = sdkerrors.Register(Codespace, 5, "invalid ServiceFeeCap")
	ErrInvalidDescription   = sdkerrors.Register(Codespace, 6, "invalid Description")
	ErrValidateTimeout      = sdkerrors.Register(Codespace, 7, "ValidateTimeout  error")
	ErrInvalidState         = sdkerrors.Register(Codespace, 8, "invalid state")
	ErrInvalidFeedName      = sdkerrors.Register(Codespace, 9, "invalid FeedName")
	ErrGenConn              = sdkerrors.Register(Codespace, 10, "generate conn error")
	ErrAccAddressFromBech32 = sdkerrors.Register(Codespace, 11, "AccAddressFromBech32 error ")
	ErrQueryFeed            = sdkerrors.Register(Codespace, 12, "query feed error ")
	ErrQueryFeedValue       = sdkerrors.Register(Codespace, 13, "query feed value error ")
	ErrMatchString          = sdkerrors.Register(Codespace, 14, "MatchString error ")
	ErrInvalidTimeout       = sdkerrors.Register(Codespace, 15, "invalid timeout")
	ErrResponseThreshold    = sdkerrors.Register(Codespace, 16, "invalid ResponseThreshold")
	ErrAggregateFunc        = sdkerrors.Register(Codespace, 17, "invalid aggregateFunc")
	ErrLatestHistory        = sdkerrors.Register(Codespace, 18, "invalid latest history")
	ErrInvalidServiceFeeCap = sdkerrors.Register(Codespace, 19, "invalid ServiceFeeCap")
)
