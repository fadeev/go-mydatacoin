package keeper

import (
	"mdc/x/mdc/types"
)

var _ types.QueryServer = Keeper{}
