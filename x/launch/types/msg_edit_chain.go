package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spntypes "github.com/tendermint/spn/pkg/types"
)

const TypeMsgEditChain = "edit_chain"

var _ sdk.Msg = &MsgEditChain{}

func NewMsgEditChain(
	coordinator string,
	launchID uint64,
	setCampaign bool,
	campaignID uint64,
	metadata []byte,
) *MsgEditChain {
	return &MsgEditChain{
		Coordinator:   coordinator,
		LaunchID:      launchID,
		SetCampaignID: setCampaign,
		CampaignID:    campaignID,
		Metadata:      metadata,
	}
}

func (msg *MsgEditChain) Route() string {
	return RouterKey
}

func (msg *MsgEditChain) Type() string {
	return TypeMsgEditChain
}

func (msg *MsgEditChain) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgEditChain) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgEditChain) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if len(msg.Metadata) == 0 && !msg.SetCampaignID {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no value to edit")
	}

	// TODO parameterize
	if len(msg.Metadata) > spntypes.MaxMetadataLength {
		return sdkerrors.Wrapf(ErrInvalidMetadataLength, "data length %d is greater than maximum %d",
			len(msg.Metadata), spntypes.MaxMetadataLength)
	}

	return nil
}
