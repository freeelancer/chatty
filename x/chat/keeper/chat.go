package keeper

import (
	"chatty/x/chat/types"
	"fmt"

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
		Message:   message,
		Encrypted: encrypted,
		CreatedAt: ctx.BlockTime().Unix(),
	}
	conversation.LastMessageAt = ctx.BlockTime().Unix()
	conversation.Messages = append(conversation.Messages, &chatMessage)
	return k.SetConversation(ctx, *conversation)
}

// CreateGroupConversation creates a group conversation
func (k Keeper) CreateGroupConversation(ctx sdk.Context, admin sdk.AccAddress, name string, participants []string, message, pubkey string) error {
	params := k.GetParams(ctx)
	convoId := params.GroupConversationCounter
	params.GroupConversationCounter = params.GroupConversationCounter + 1
	if err := k.SetParams(ctx, params); err != nil {
		return err
	}

	for _, participant := range participants {
		if err := k.AddIdToAddressGroup(ctx, participant, convoId); err != nil {
			return err
		}
	}

	groupConvo := types.GroupConversation{
		Id:            convoId,
		Admin:         admin.String(),
		Name:          name,
		Participants:  participants,
		Messages:      []*types.ChatMessage{},
		LastMessageAt: ctx.BlockTime().Unix(),
		CreatedAt:     ctx.BlockTime().Unix(),
	}

	encrypted := false
	if pubkey != "" {
		encrypted = true
		groupConvo.PubKey = &types.PubKey{
			Address: admin.String(),
			Key:     pubkey,
		}
	}

	if message != "" {
		chatMessage := types.ChatMessage{
			Sender:    admin.String(),
			Message:   message,
			Encrypted: encrypted,
			CreatedAt: ctx.BlockTime().Unix(),
		}
		groupConvo.Messages = append(groupConvo.Messages, &chatMessage)
	}

	return k.SetGroupConversation(ctx, groupConvo)
}

// CreateGroupConversationMessage for Group Conversation
func (k Keeper) CreateGroupConversationMessage(ctx sdk.Context, sender sdk.AccAddress, id uint64, message string, encrypted bool) error {
	groupConvo, hasGroupConvo := k.GetGroupConversation(ctx, id)
	if !hasGroupConvo {
		return fmt.Errorf("group conversation with id %d does not exist", id)
	}

	canSend := false
	for _, participant := range groupConvo.Participants {
		if participant == sender.String() {
			canSend = true
			break
		}
	}
	if !canSend {
		return fmt.Errorf("sender %s is not a participant of group conversation with id %d", sender.String(), id)
	}

	chatMessage := types.ChatMessage{
		Sender:    sender.String(),
		Message:   message,
		Encrypted: encrypted,
		CreatedAt: ctx.BlockTime().Unix(),
	}
	groupConvo.LastMessageAt = ctx.BlockTime().Unix()
	groupConvo.Messages = append(groupConvo.Messages, &chatMessage)
	return k.SetGroupConversation(ctx, *groupConvo)
}

func (k Keeper) UpdatePubkey(ctx sdk.Context, creator string, pubkeyStr string) error {
	pubkey := types.PubKey{
		Address: creator,
		Key:     pubkeyStr,
	}
	return k.SetPubkey(ctx, pubkey)
}
