package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgProposeSend = "propose_send"

var _ sdk.Msg = &MsgProposeSend{}

func NewMsgProposeSend(creator string, title string, description string, beneficiary string, coins sdk.Coins) *MsgProposeSend {
	return &MsgProposeSend{
		Creator:     creator,
		Title:       title,
		Description: description,
		Beneficiary: beneficiary,
		Coins:       coins,
	}
}

func (msg *MsgProposeSend) Route() string {
	return RouterKey
}

func (msg *MsgProposeSend) Type() string {
	return TypeMsgProposeSend
}

func (msg *MsgProposeSend) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgProposeSend) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgProposeSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
