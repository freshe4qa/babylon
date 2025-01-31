package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/zoneconcierge module sentinel errors
var (
	ErrSample                  = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout    = sdkerrors.Register(ModuleName, 1101, "invalid packet timeout")
	ErrInvalidVersion          = sdkerrors.Register(ModuleName, 1102, "invalid version")
	ErrHeaderNotFound          = sdkerrors.Register(ModuleName, 1103, "no header exists at this height")
	ErrInvalidHeader           = sdkerrors.Register(ModuleName, 1104, "input header is invalid")
	ErrNoValidAncestorHeader   = sdkerrors.Register(ModuleName, 1105, "no valid ancestor for this header")
	ErrForkNotFound            = sdkerrors.Register(ModuleName, 1106, "cannot find fork")
	ErrInvalidForks            = sdkerrors.Register(ModuleName, 1107, "input forks is invalid")
	ErrChainInfoNotFound       = sdkerrors.Register(ModuleName, 1108, "no chain info exists")
	ErrEpochChainInfoNotFound  = sdkerrors.Register(ModuleName, 1109, "no chain info exists at this epoch")
	ErrEpochHeadersNotFound    = sdkerrors.Register(ModuleName, 1110, "no timestamped header exists at this epoch")
	ErrFinalizedEpochNotFound  = sdkerrors.Register(ModuleName, 1111, "cannot find a finalized epoch")
	ErrInvalidProofEpochSealed = sdkerrors.Register(ModuleName, 1112, "invalid ProofEpochSealed")
	ErrInvalidMerkleProof      = sdkerrors.Register(ModuleName, 1113, "invalid Merkle inclusion proof")
	ErrInvalidChainInfo        = sdkerrors.Register(ModuleName, 1114, "invalid chain info")
)
