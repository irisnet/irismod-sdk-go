package nft

import (
	sdkerrors "github.com/irisnet/core-sdk-go/types"
)

var (
	ErrInvalidCollection = sdkerrors.RegisterErr(ModuleName, 2, "invalid nft collection")
	ErrUnknownCollection = sdkerrors.RegisterErr(ModuleName, 3, "unknown nft collection")
	ErrInvalidNFT        = sdkerrors.RegisterErr(ModuleName, 4, "invalid nft")
	ErrNFTAlreadyExists  = sdkerrors.RegisterErr(ModuleName, 5, "nft already exists")
	ErrUnknownNFT        = sdkerrors.RegisterErr(ModuleName, 6, "unknown nft")
	ErrEmptyTokenData    = sdkerrors.RegisterErr(ModuleName, 7, "nft data can't be empty")
	ErrUnauthorized      = sdkerrors.RegisterErr(ModuleName, 8, "unauthorized address")
	ErrInvalidDenom      = sdkerrors.RegisterErr(ModuleName, 9, "invalid denom")
	ErrInvalidTokenID    = sdkerrors.RegisterErr(ModuleName, 10, "invalid nft id")
	ErrInvalidTokenURI   = sdkerrors.RegisterErr(ModuleName, 11, "invalid nft uri")
)
