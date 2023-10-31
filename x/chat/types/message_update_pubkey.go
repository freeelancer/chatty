package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdatePubkey = "update_pubkey"

var _ sdk.Msg = &MsgUpdatePubkey{}

func NewMsgUpdatePubkey(creator, pubkeyStr string) *MsgUpdatePubkey {
	return &MsgUpdatePubkey{
		Creator: creator,
		Pubkey:  pubkeyStr,
	}
}

func (msg *MsgUpdatePubkey) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePubkey) Type() string {
	return TypeMsgUpdatePubkey
}

func (msg *MsgUpdatePubkey) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePubkey) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePubkey) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
