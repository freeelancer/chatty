package keeper

import (
	"context"

	"chatty/x/chat/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateChatMessage(goCtx context.Context, msg *types.MsgCreateChatMessage) (*types.MsgCreateChatMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateChatMessageResponse{}, nil
}
