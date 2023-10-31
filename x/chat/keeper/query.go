package keeper

import (
	"chatty/x/chat/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Conversation(goCtx context.Context, req *types.QueryConversationRequest) (*types.QueryConversationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	conversations := []*types.Conversation{}
	if req.ChatterB == "" {
		store := ctx.KVStore(k.storeKey)
		iter := sdk.KVStorePrefixIterator(store, append(types.ConversationKeyPrefix, []byte(req.ChatterA)...))
		defer iter.Close()
		for ; iter.Valid(); iter.Next() {
			var conversation types.Conversation
			k.cdc.MustUnmarshal(iter.Value(), &conversation)
			conversations = append(conversations, &conversation)
		}
	} else {
		conversation, hasConversation := k.GetConversation(ctx, req.ChatterA, req.ChatterB)
		if hasConversation {
			conversations = append(conversations, conversation)
		}
	}

	return &types.QueryConversationResponse{Conversations: conversations}, nil
}
