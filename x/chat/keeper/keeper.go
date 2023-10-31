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
	store.Set(types.GetAddressPubkeyKey(pubkey.Address), k.cdc.MustMarshal(&pubkey))
	return nil
}

// GetPubkey gets a pubkey from the store.
func (k Keeper) GetPubkey(ctx sdk.Context, address string) (*types.PubKey, bool) {
	store := ctx.KVStore(k.storeKey)
	pubkeyBs := store.Get(types.GetAddressPubkeyKey(address))
	if pubkeyBs == nil {
		return nil, false
	}
	pubkey := types.PubKey{}
	k.cdc.MustUnmarshal(pubkeyBs, &pubkey)
	return &pubkey, true
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

// SetGroupConversation sets a group conversation to the store.
func (k Keeper) SetGroupConversation(ctx sdk.Context, groupConversation types.GroupConversation) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetGroupConversationKey(groupConversation.Id), k.cdc.MustMarshal(&groupConversation))
	return nil
}

// GetGroupConversation gets a group conversation from the store.
func (k Keeper) GetGroupConversation(ctx sdk.Context, id int64) (*types.GroupConversation, bool) {
	store := ctx.KVStore(k.storeKey)
	groupConversationBs := store.Get(types.GetGroupConversationKey(id))
	if groupConversationBs == nil {
		return nil, false
	}
	groupConversation := types.GroupConversation{}
	k.cdc.MustUnmarshal(groupConversationBs, &groupConversation)
	return &groupConversation, true
}

// GetAddressGroup gets a group conversation ids from an address.
func (k Keeper) GetAddressGroup(ctx sdk.Context, address string) (*types.AddressGroups, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetAddressGroupKey(address)
	idsBs := store.Get(key)
	if idsBs == nil {
		return nil, false
	}
	addressGroups := types.AddressGroups{}
	k.cdc.MustUnmarshal(idsBs, &addressGroups)
	return &addressGroups, true
}

// AddIdToAddressGroup adds a group conversation id to an address.
func (k Keeper) AddIdToAddressGroup(ctx sdk.Context, address string, id int64) error {
	store := ctx.KVStore(k.storeKey)
	key := types.GetAddressGroupKey(address)
	addressGroup, found := k.GetAddressGroup(ctx, address)
	if !found {
		addressGroup = &types.AddressGroups{
			Address:  address,
			GroupIds: []int64{},
		}
	}
	addressGroup.GroupIds = append(addressGroup.GroupIds, id)
	store.Set(key, k.cdc.MustMarshal(addressGroup))
	return nil
}
