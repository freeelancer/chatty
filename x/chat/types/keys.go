package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "chat"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_chat"
)

var (
	ParamsKey                  = []byte("p_chat")
	ConversationKeyPrefix      = []byte{0x01}
	AddressPubkeyKeyPrefix     = []byte{0x02}
	GroupConversationKeyPrefix = []byte{0x03}
	AddressGroupKeyPrefix      = []byte{0x04}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetConversationKey(sender, receiver string) ([]byte, []byte) {
	return append(ConversationKeyPrefix, []byte(sender+receiver)...), append(ConversationKeyPrefix, []byte(receiver+sender)...)
}

func GetAddressPubkeyKey(creator string) []byte {
	return append(AddressPubkeyKeyPrefix, []byte(creator)...)
}

func GetGroupConversationKey(id int64) []byte {
	return append(GroupConversationKeyPrefix, sdk.Uint64ToBigEndian(uint64(id))...)
}

func GetAddressGroupKey(address string) []byte {
	return append(AddressGroupKeyPrefix, []byte(address)...)
}
