package htlc

import (
	sdkerrors "github.com/irisnet/core-sdk-go/types"
)

// HTLC module sentinel errors
var (
	ErrInvalidID                   = sdkerrors.RegisterErr(ModuleName, 2, "invalid htlc id")
	ErrInvalidHashLock             = sdkerrors.RegisterErr(ModuleName, 3, "invalid hash lock")
	ErrInvalidTimeLock             = sdkerrors.RegisterErr(ModuleName, 4, "invalid time lock")
	ErrInvalidSecret               = sdkerrors.RegisterErr(ModuleName, 5, "invalid secret")
	ErrInvalidExpirationHeight     = sdkerrors.RegisterErr(ModuleName, 6, "invalid expiration height")
	ErrInvalidTimestamp            = sdkerrors.RegisterErr(ModuleName, 7, "invalid timestamp")
	ErrInvalidState                = sdkerrors.RegisterErr(ModuleName, 8, "invalid state")
	ErrInvalidClosedBlock          = sdkerrors.RegisterErr(ModuleName, 9, "invalid closed block")
	ErrInvalidDirection            = sdkerrors.RegisterErr(ModuleName, 10, "invalid direction")
	ErrHTLCExists                  = sdkerrors.RegisterErr(ModuleName, 11, "htlc already exists")
	ErrUnknownHTLC                 = sdkerrors.RegisterErr(ModuleName, 12, "unknown htlc")
	ErrHTLCNotOpen                 = sdkerrors.RegisterErr(ModuleName, 13, "htlc not open")
	ErrAssetNotSupported           = sdkerrors.RegisterErr(ModuleName, 14, "asset not found")
	ErrAssetNotActive              = sdkerrors.RegisterErr(ModuleName, 15, "asset is currently inactive")
	ErrInvalidAccount              = sdkerrors.RegisterErr(ModuleName, 16, "invalid account")
	ErrInvalidAmount               = sdkerrors.RegisterErr(ModuleName, 17, "invalid amount")
	ErrInsufficientAmount          = sdkerrors.RegisterErr(ModuleName, 18, "amount cannot cover the deputy fixed fee")
	ErrExceedsSupplyLimit          = sdkerrors.RegisterErr(ModuleName, 19, "asset supply over limit")
	ErrExceedsTimeBasedSupplyLimit = sdkerrors.RegisterErr(ModuleName, 20, "asset supply over limit for current time period")
	ErrInvalidCurrentSupply        = sdkerrors.RegisterErr(ModuleName, 21, "supply decrease puts current asset supply below 0")
	ErrInvalidIncomingSupply       = sdkerrors.RegisterErr(ModuleName, 22, "supply decrease puts incoming asset supply below 0")
	ErrInvalidOutgoingSupply       = sdkerrors.RegisterErr(ModuleName, 23, "supply decrease puts outgoing asset supply below 0")
	ErrExceedsAvailableSupply      = sdkerrors.RegisterErr(ModuleName, 24, "outgoing swap exceeds total available supply")
	ErrAssetSupplyNotFound         = sdkerrors.RegisterErr(ModuleName, 25, "asset supply not found in store")
)
