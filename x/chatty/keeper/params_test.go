package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "chatty/testutil/keeper"
	"chatty/x/chatty/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.ChattyKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
