package service

import sdkerrors "github.com/irisnet/core-sdk-go/types/errors"

const Codespace = sdkerrors.RootCodespace + ModuleName

var (
	ErrQueryAddress                 = sdkerrors.Register(Codespace, 1, "query address error")
	ErrToMintCoin                   = sdkerrors.Register(Codespace, 2, " to mint coin error")
	ErrReqCtxID                     = sdkerrors.Register(Codespace, 3, " length error")
	ErrGenConn                      = sdkerrors.Register(Codespace, 4, "generate conn error")
	ErrQueryParams                  = sdkerrors.Register(Codespace, 5, "query params error")
	ErrValidateAccAddress           = sdkerrors.Register(Codespace, 6, "ValidateAccAddress error ")
	ErrBuildAndSend                 = sdkerrors.Register(Codespace, 7, "BuildAndSend error ")
	ErrEventsGetValue               = sdkerrors.Register(Codespace, 8, "Events.GetValue error")
	ErrSubscribeServiceResponse     = sdkerrors.Register(Codespace, 9, "SubscribeServiceResponse error ")
	ErrQueryDefinition              = sdkerrors.Register(Codespace, 10, "QueryDefinition  error ")
	ErrQueryBunding                 = sdkerrors.Register(Codespace, 11, "QueryBunding error ")
	ErrQueryRequestByTxQuery        = sdkerrors.Register(Codespace, 12, "QueryRequestByTxQuery error ")
	ErrQueryRequests                = sdkerrors.Register(Codespace, 13, "QueryRequests error ")
	ErrQueryResponseByTxQuery       = sdkerrors.Register(Codespace, 14, "QueryResponseByTxQuery error ")
	ErrQueryResponses               = sdkerrors.Register(Codespace, 15, "QueryResponses error ")
	ErrQueryEarnedFees              = sdkerrors.Register(Codespace, 16, "QueryEarnedFees error ")
	ErrQueryRequestContextByTxQuery = sdkerrors.Register(Codespace, 17, "QueryRequestContextByTxQuery error ")
)
