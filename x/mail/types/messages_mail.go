package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgSend = "create_mail"
)

var _ sdk.Msg = &MsgSend{}

func NewMsgSend(sender string, to string, subject string, body string) *MsgSend {
	return &MsgSend{
		Sender:  sender,
		To:      to,
		Subject: subject,
		Body:    body,
	}
}

func (msg *MsgSend) Route() string {
	return RouterKey
}

func (msg *MsgSend) Type() string {
	return TypeMsgSend
}

func (msg *MsgSend) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgSend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}
	return nil
}
