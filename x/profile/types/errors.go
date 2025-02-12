package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/profile module sentinel errors
var (
	ErrCoordAlreadyExist    = sdkerrors.Register(ModuleName, 2, "coordinator address already exist")
	ErrCoordAddressNotFound = sdkerrors.Register(ModuleName, 3, "coordinator address not found")
	ErrCoordInvalid         = sdkerrors.Register(ModuleName, 4, "invalid coordinator")
	ErrEmptyDescription     = sdkerrors.Register(ModuleName, 5, "you must provide at least one description parameter")
	ErrDupAddress           = sdkerrors.Register(ModuleName, 6, "address is duplicated")
	ErrCoordInactive        = sdkerrors.Register(ModuleName, 7, "inactive coordinator")
)
