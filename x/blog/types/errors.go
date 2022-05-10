package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/blog module sentinel errors
var (
	ErrID         = sdkerrors.Register(ModuleName, 1100, "Post ID Not Exists For Comment")
	ErrCommentOld = sdkerrors.Register(ModuleName, 1200, "More Than 100 Blocks Passed From The Post")
)
