package keeper

import (
	"chatty/x/chat/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateChatMessage(goCtx context.Context, msg *types.MsgCreateChatMessage) (*types.MsgCreateChatMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	receiver, err := sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return nil, err
	}
	if err := k.Keeper.CreateChatMessage(ctx, creator, receiver, msg.Message, msg.Encrypted); err != nil {
		return nil, err
	}

	return &types.MsgCreateChatMessageResponse{}, nil
}

func (k msgServer) UpdatePubkey(goCtx context.Context, msg *types.MsgUpdatePubkey) (*types.MsgUpdatePubkeyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.UpdatePubkey(ctx, msg.Creator, msg.Pubkey); err != nil {
		return nil, err
	}

	return &types.MsgUpdatePubkeyResponse{}, nil
}
