package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTriggerLaunch = "trigger_launch"

var _ sdk.Msg = &MsgTriggerLaunch{}

func NewMsgTriggerLaunch(coordinator string, launchID uint64, remainingTime int64) *MsgTriggerLaunch {
	return &MsgTriggerLaunch{
		Coordinator:   coordinator,
		LaunchID:      launchID,
		RemainingTime: remainingTime,
	}
}

func (msg *MsgTriggerLaunch) Route() string {
	return RouterKey
}

func (msg *MsgTriggerLaunch) Type() string {
	return TypeMsgTriggerLaunch
}

func (msg *MsgTriggerLaunch) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgTriggerLaunch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTriggerLaunch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid coordinator address (%s)", err)
	}

	if msg.RemainingTime <= 0 {
		return sdkerrors.Wrapf(ErrRemainingTimeNotPositive, "value must be greater than 0, %v <= 0", msg.RemainingTime)
	}
	return nil
}
