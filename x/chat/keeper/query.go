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

	conversation, hasConversation := k.GetConversation(ctx, req.ChatterA, req.ChatterB)
	if !hasConversation {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryConversationResponse{Conversation: conversation}, nil
}

func (k Keeper) Conversations(goCtx context.Context, req *types.QueryConversationsRequest) (*types.QueryConversationsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	conversations := []*types.Conversation{}
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, types.ConversationKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		conversation := &types.Conversation{}
		k.cdc.MustUnmarshal(iter.Value(), conversation)
		conversations = append(conversations, conversation)
	}

	return &types.QueryConversationsResponse{Conversations: conversations}, nil
}
