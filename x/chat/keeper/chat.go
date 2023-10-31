package keeper

import (
	"chatty/x/chat/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CreateChatMessage(ctx sdk.Context, sender, receiver sdk.AccAddress, message string, encrypted bool) error {
	conversation, hasConversation := k.GetConversation(ctx, sender.String(), receiver.String())
	if !hasConversation {
		conversation = &types.Conversation{
			ChatterA:  sender.String(),
			ChatterB:  receiver.String(),
			CreatedAt: ctx.BlockTime().Unix(),
		}
	}
	chatMessage := types.ChatMessage{
		Sender:    sender.String(),
		Receiver:  receiver.String(),
		Message:   message,
		Encrypted: encrypted,
		CreatedAt: ctx.BlockTime().Unix(),
	}
	conversation.LastMessageAt = ctx.BlockTime().Unix()
	conversation.Messages = append(conversation.Messages, &chatMessage)
	return k.SetConversation(ctx, *conversation)
}

func (k Keeper) UpdatePubkey(ctx sdk.Context, creator string, pubkeyStr string) error {
	pubkey := types.PubKey{
		Address: creator,
		Key:     pubkeyStr,
	}
	return k.SetPubkey(ctx, pubkey)
}
