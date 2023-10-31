package keeper

import (
	"chatty/x/chat/types"
)

var _ types.QueryServer = Keeper{}
