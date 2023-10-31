package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateChatMessage = "create_chat_message"

var _ sdk.Msg = &MsgCreateChatMessage{}

func NewMsgCreateChatMessage(creator, receiver, message string, encrypted bool) *MsgCreateChatMessage {
	return &MsgCreateChatMessage{
		Creator:   creator,
		Receiver:  receiver,
		Message:   message,
		Encrypted: encrypted,
	}
}

func (msg *MsgCreateChatMessage) Route() string {
	return RouterKey
}

func (msg *MsgCreateChatMessage) Type() string {
	return TypeMsgCreateChatMessage
}

func (msg *MsgCreateChatMessage) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateChatMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateChatMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address (%s)", err)
	}
	if msg.Creator == msg.Receiver {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "creator and receiver are the same")
	}
	if msg.Message == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "message is empty")
	}
	return nil
}
