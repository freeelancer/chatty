package chat

import (
	"chatty/x/chat/keeper"
	"chatty/x/chat/types"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for _, pubkey := range genState.PubKeys {
		if err := k.SetPubkey(ctx, *pubkey); err != nil {
			panic(err)
		}
	}

	for _, conversation := range genState.Conversations {
		if err := k.SetConversation(ctx, *conversation); err != nil {
			panic(err)
		}
	}

	for _, addressGroup := range genState.AddressGroups {
		if err := k.SetAddressGroup(ctx, *addressGroup); err != nil {
			panic(err)
		}
	}

	for _, groupConversation := range genState.GroupConversations {
		if err := k.SetGroupConversation(ctx, *groupConversation); err != nil {
			panic(err)
		}
	}
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	pubkeysResp, err := k.Pubkeys(ctx.Context(), &types.QueryPubkeysRequest{})
	if err != nil {
		panic(err)
	}
	genesis.PubKeys = pubkeysResp.Pubkeys

	conversationsResp, err := k.Conversations(ctx, &types.QueryConversationsRequest{})
	if err != nil {
		panic(err)
	}
	genesis.Conversations = conversationsResp.Conversations

	addressGroups := k.GetAllAddressGroup(ctx)
	genesis.AddressGroups = addressGroups

	var groupConversations []*types.GroupConversation
	for id := int64(1); id < genesis.Params.GroupConversationCounter; id++ {
		groupConversation, found := k.GetGroupConversation(ctx, id)
		if !found {
			panic(fmt.Sprintf("group conversation %d not found", id))
		}
		groupConversations = append(genesis.GroupConversations, groupConversation)
	}
	genesis.GroupConversations = groupConversations

	// this line is used by starport scaffolding # genesis/module/export
	return genesis
}
