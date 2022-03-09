package keeper

import (
	"github.com/notional-labs/tinyportapp/x/tinyportapp/types"
)

var _ types.QueryServer = Keeper{}
