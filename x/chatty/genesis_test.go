package chatty_test

import (
	"testing"

	keepertest "chatty/testutil/keeper"
	"chatty/testutil/nullify"
	"chatty/x/chatty"
	"chatty/x/chatty/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChattyKeeper(t)
	chatty.InitGenesis(ctx, k, genesisState)
	got := chatty.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
