package keeper

import (
	"github.com/mohammadreza-torkaman/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
