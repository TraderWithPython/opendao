package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/od module sentinel errors
var (
	ErrSample            = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrNotMP             = sdkerrors.Register(ModuleName, 1101, "not mp")
	ErrProposalNotFound  = sdkerrors.Register(ModuleName, 1102, "proposal not found")
	ErrInsufficientStake = sdkerrors.Register(ModuleName, 1103, "insufficient stake")
)
