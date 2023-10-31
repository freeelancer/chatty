package types

import (
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGroupConversation = "create_group_conversation"

var _ sdk.Msg = &MsgCreateGroupConversation{}

func NewMsgCreateGroupConversation(creator, name, message, pubkey string, participants []string, encrypted bool) *MsgCreateGroupConversation {
	return &MsgCreateGroupConversation{
		Creator:      creator,
		Name:         name,
		Participants: participants,
		Message:      message,
		Pubkey:       pubkey,
		Encrypted:    encrypted,
	}
}

func (msg *MsgCreateGroupConversation) Route() string {
	return RouterKey
}

func (msg *MsgCreateGroupConversation) Type() string {
	return TypeMsgCreateGroupConversation
}

func (msg *MsgCreateGroupConversation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGroupConversation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGroupConversation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Pubkey != "" {
		keyBytes, err := hex.DecodeString(msg.Pubkey)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid pubkey (%s)", err)
		}
		block, _ := pem.Decode(keyBytes)
		if block == nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Error decoding PEM block")
		}

		if block.Type != "RSA PUBLIC KEY" {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Found block of type %s instead of RSA PUBLIC KEY", block.Type)
		}
		// Parse the RSA public key
		_, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Error parsing RSA public key:", err)
		}
	}

	if msg.Name == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "name cannot be empty")
	}

	if msg.Participants == nil || len(msg.Participants) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "participants cannot be empty")
	}

	if len(msg.Participants) == 1 {
		if msg.Participants[0] == msg.Creator {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "participants only has creator")
		}
	}

	for _, participant := range msg.Participants {
		_, err := sdk.AccAddressFromBech32(participant)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid participant address (%s)", err)
		}
	}

	return nil
}
