package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDisableCoordinator = "disable_coordinator"

var _ sdk.Msg = &MsgDisableCoordinator{}

func NewMsgDisableCoordinator(address string) *MsgDisableCoordinator {
	return &MsgDisableCoordinator{
		Address: address,
	}
}

func (msg *MsgDisableCoordinator) Route() string {
	return RouterKey
}

func (msg *MsgDisableCoordinator) Type() string {
	return TypeMsgDisableCoordinator
}

func (msg *MsgDisableCoordinator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDisableCoordinator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDisableCoordinator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
