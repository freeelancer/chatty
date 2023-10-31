package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateChatMessage = "create_chat_message"

var _ sdk.Msg = &MsgCreateChatMessage{}

func NewMsgCreateChatMessage(creator string) *MsgCreateChatMessage {
	return &MsgCreateChatMessage{
		Creator: creator,
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
	return nil
}
