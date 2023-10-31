package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"chatty/x/chat/types"
)

type (
	Keeper struct {
		cdc      codec.BinaryCodec
		storeKey storetypes.StoreKey
		memKey   storetypes.StoreKey

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	authority string,

) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		memKey:    memKey,
		authority: authority,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetPubkey sets a pubkey to the store.
func (k Keeper) SetPubkey(ctx sdk.Context, pubkey types.PubKey) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetPubkeyKey(pubkey.Address), k.cdc.MustMarshal(&pubkey))
	return nil
}

// SetConversation sets a conversation to the store.
func (k Keeper) SetConversation(ctx sdk.Context, conversation types.Conversation) error {
	store := ctx.KVStore(k.storeKey)
	key1, key2 := types.GetConversationKey(conversation.ChatterA, conversation.ChatterB)
	store.Set(key1, k.cdc.MustMarshal(&conversation))
	store.Set(key2, k.cdc.MustMarshal(&conversation))
	return nil
}

// GetConversation gets a conversation from the store.
func (k Keeper) GetConversation(ctx sdk.Context, chatterA, chatterB string) (*types.Conversation, bool) {
	store := ctx.KVStore(k.storeKey)
	key, _ := types.GetConversationKey(chatterA, chatterB)
	conversationBs := store.Get(key)
	if conversationBs == nil {
		return nil, false
	}
	conversation := types.Conversation{}
	k.cdc.MustUnmarshal(conversationBs, &conversation)
	return &conversation, true
}

// SetChatMessage sets a chat message to the store.
func (k Keeper) SetChatMessage(ctx sdk.Context, chatMessage types.ChatMessage) error {
	conversation, hasConversation := k.GetConversation(ctx, chatMessage.Sender, chatMessage.Receiver)
	if hasConversation {
		return fmt.Errorf("conversation not found")
	}
	conversation.Messages = append(conversation.Messages, &chatMessage)
	return k.SetConversation(ctx, *conversation)
}
