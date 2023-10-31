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

func (k Keeper) Pubkey(goCtx context.Context, req *types.QueryPubkeyRequest) (*types.QueryPubkeyResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pubkey, hasPubkey := k.GetPubkey(ctx, req.Address)
	if !hasPubkey {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryPubkeyResponse{Pubkey: pubkey}, nil
}

func (k Keeper) Pubkeys(goCtx context.Context, req *types.QueryPubkeysRequest) (*types.QueryPubkeysResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	store := ctx.KVStore(k.storeKey)
	pubkeys := []*types.PubKey{}
	iter := sdk.KVStorePrefixIterator(store, types.PubkeyKeyPrefix)
	defer iter.Close()
	for ; iter.Valid(); iter.Next() {
		pubkey := &types.PubKey{}
		k.cdc.MustUnmarshal(iter.Value(), pubkey)
		pubkeys = append(pubkeys, pubkey)
	}

	return &types.QueryPubkeysResponse{Pubkeys: pubkeys}, nil
}
