package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGroupConversationMessage = "create_group_conversation_message"

var _ sdk.Msg = &MsgCreateGroupConversationMessage{}

func NewMsgCreateGroupConversationMessage(creator, message string, conversationId int64, encrypted bool) *MsgCreateGroupConversationMessage {
	return &MsgCreateGroupConversationMessage{
		Creator:        creator,
		ConversationId: conversationId,
		Messsage:       message,
		Encrypted:      encrypted,
	}
}

func (msg *MsgCreateGroupConversationMessage) Route() string {
	return RouterKey
}

func (msg *MsgCreateGroupConversationMessage) Type() string {
	return TypeMsgCreateGroupConversationMessage
}

func (msg *MsgCreateGroupConversationMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGroupConversationMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGroupConversationMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
