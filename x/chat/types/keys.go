package types

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
	ParamsKey             = []byte("p_chat")
	ConversationKeyPrefix = []byte{0x01}
	PubkeyKeyPrefix       = []byte{0x02}
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func GetConversationKey(sender, receiver string) ([]byte, []byte) {
	return append(ConversationKeyPrefix, []byte(sender+receiver)...), append(ConversationKeyPrefix, []byte(receiver+sender)...)
}

func GetPubkeyKey(creator string) []byte {
	return append(PubkeyKeyPrefix, []byte(creator)...)
}
