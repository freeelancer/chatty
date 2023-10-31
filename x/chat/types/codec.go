package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	// this line is used by starport scaffolding # 1
)

// RegisterLegacyAminoCodec registers concrete types on the LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateChatMessage{}, "chat/CreateChatMessage", nil)
	cdc.RegisterConcrete(&MsgUpdatePubkey{}, "chat/UpdatePubkey", nil)
	cdc.RegisterConcrete(&MsgCreateGroupConversation{}, "chat/CreateGroupConversation", nil)
	cdc.RegisterConcrete(&MsgCreateGroupConversationMessage{}, "chat/CreateGroupConversationMessage", nil)
	// this line is used by starport scaffolding # 2

	cdc.RegisterConcrete(Params{}, "chatty/x/chat/Params", nil)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, "chatty/x/chat/MsgUpdateParams")
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChatMessage{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdatePubkey{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroupConversation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateGroupConversationMessage{},
	)
	// this line is used by starport scaffolding # 3

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateParams{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
