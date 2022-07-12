package service

import (
	sdkerrors "github.com/irisnet/core-sdk-go/types"
)

// service module sentinel errors
var (
	ErrInvalidServiceName       = sdkerrors.RegisterErr(ModuleName, 2, "invalid service name, must contain alphanumeric characters, _ and - only，length greater than 0 and less than or equal to 70")
	ErrInvalidDescription       = sdkerrors.RegisterErr(ModuleName, 3, "invalid description")
	ErrInvalidTags              = sdkerrors.RegisterErr(ModuleName, 4, "invalid tags")
	ErrInvalidSchemas           = sdkerrors.RegisterErr(ModuleName, 5, "invalid schemas")
	ErrUnknownServiceDefinition = sdkerrors.RegisterErr(ModuleName, 6, "unknown service definition")
	ErrServiceDefinitionExists  = sdkerrors.RegisterErr(ModuleName, 7, "service definition already exists")

	ErrInvalidDeposit            = sdkerrors.RegisterErr(ModuleName, 8, "invalid deposit")
	ErrInvalidMinDeposit         = sdkerrors.RegisterErr(ModuleName, 9, "invalid minimum deposit")
	ErrInvalidPricing            = sdkerrors.RegisterErr(ModuleName, 10, "invalid pricing")
	ErrInvalidQoS                = sdkerrors.RegisterErr(ModuleName, 11, "invalid QoS")
	ErrInvalidOptions            = sdkerrors.RegisterErr(ModuleName, 12, "invalid options")
	ErrServiceBindingExists      = sdkerrors.RegisterErr(ModuleName, 13, "service binding already exists")
	ErrUnknownServiceBinding     = sdkerrors.RegisterErr(ModuleName, 14, "unknown service binding")
	ErrServiceBindingUnavailable = sdkerrors.RegisterErr(ModuleName, 15, "service binding unavailable")
	ErrServiceBindingAvailable   = sdkerrors.RegisterErr(ModuleName, 16, "service binding available")
	ErrIncorrectRefundTime       = sdkerrors.RegisterErr(ModuleName, 17, "incorrect refund time")

	ErrInvalidServiceFeeCap      = sdkerrors.RegisterErr(ModuleName, 18, "invalid service fee cap")
	ErrInvalidProviders          = sdkerrors.RegisterErr(ModuleName, 19, "invalid providers")
	ErrInvalidTimeout            = sdkerrors.RegisterErr(ModuleName, 20, "invalid timeout")
	ErrInvalidRepeatedFreq       = sdkerrors.RegisterErr(ModuleName, 21, "invalid repeated frequency")
	ErrInvalidRepeatedTotal      = sdkerrors.RegisterErr(ModuleName, 22, "invalid repeated total count")
	ErrInvalidResponseThreshold  = sdkerrors.RegisterErr(ModuleName, 23, "invalid response threshold")
	ErrInvalidResponse           = sdkerrors.RegisterErr(ModuleName, 24, "invalid response")
	ErrInvalidRequestID          = sdkerrors.RegisterErr(ModuleName, 25, "invalid request ID")
	ErrUnknownRequest            = sdkerrors.RegisterErr(ModuleName, 26, "unknown request")
	ErrUnknownResponse           = sdkerrors.RegisterErr(ModuleName, 27, "unknown response")
	ErrUnknownRequestContext     = sdkerrors.RegisterErr(ModuleName, 28, "unknown request context")
	ErrInvalidRequestContextID   = sdkerrors.RegisterErr(ModuleName, 29, "invalid request context ID")
	ErrRequestContextNonRepeated = sdkerrors.RegisterErr(ModuleName, 30, "request context non repeated")
	ErrRequestContextNotRunning  = sdkerrors.RegisterErr(ModuleName, 31, "request context not running")
	ErrRequestContextNotPaused   = sdkerrors.RegisterErr(ModuleName, 32, "request context not paused")
	ErrRequestContextCompleted   = sdkerrors.RegisterErr(ModuleName, 33, "request context completed")
	ErrCallbackRegistered        = sdkerrors.RegisterErr(ModuleName, 34, "callback registered")
	ErrCallbackNotRegistered     = sdkerrors.RegisterErr(ModuleName, 35, "callback not registered")
	ErrNoEarnedFees              = sdkerrors.RegisterErr(ModuleName, 36, "no earned fees")

	ErrInvalidRequestInput   = sdkerrors.RegisterErr(ModuleName, 37, "invalid request input")
	ErrInvalidResponseOutput = sdkerrors.RegisterErr(ModuleName, 38, "invalid response output")
	ErrInvalidResponseResult = sdkerrors.RegisterErr(ModuleName, 39, "invalid response result")

	ErrInvalidSchemaName = sdkerrors.RegisterErr(ModuleName, 40, "invalid service schema name")
	ErrNotAuthorized     = sdkerrors.RegisterErr(ModuleName, 41, "not authorized")

	ErrModuleServiceRegistered   = sdkerrors.RegisterErr(ModuleName, 42, "module service registered")
	ErrInvalidModuleService      = sdkerrors.RegisterErr(ModuleName, 43, "invalid module service")
	ErrBindModuleService         = sdkerrors.RegisterErr(ModuleName, 44, "can not bind module service")
	ErrInvalidRequestInputBody   = sdkerrors.RegisterErr(ModuleName, 45, "invalid request input body")
	ErrInvalidResponseOutputBody = sdkerrors.RegisterErr(ModuleName, 46, "invalid response output body")
)
