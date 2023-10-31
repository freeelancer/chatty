package keeper

import (
	"chatty/x/chatty/types"
)

var _ types.QueryServer = Keeper{}
