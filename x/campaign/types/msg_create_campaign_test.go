package types_test

import (
	"testing"

	spntypes "github.com/tendermint/spn/pkg/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestMsgCreateCampaign_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgCreateCampaign
		err  error
	}{
		{
			name: "valid message",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(r),
				CampaignName: sample.CampaignName(r),
				TotalSupply:  sample.TotalSupply(r),
				Metadata:     sample.Metadata(r, 20),
			},
		},
		{
			name: "invalid address",
			msg: types.MsgCreateCampaign{
				Coordinator:  "invalid_address",
				CampaignName: sample.CampaignName(r),
				TotalSupply:  sample.TotalSupply(r),
				Metadata:     sample.Metadata(r, 20),
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid campaign name",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(r),
				CampaignName: invalidCampaignName,
				TotalSupply:  sample.TotalSupply(r),
				Metadata:     sample.Metadata(r, 20),
			},
			err: types.ErrInvalidCampaignName,
		},
		{
			name: "invalid total supply",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(r),
				CampaignName: sample.CampaignName(r),
				TotalSupply:  invalidCoins,
				Metadata:     sample.Metadata(r, 20),
			},
			err: types.ErrInvalidTotalSupply,
		},
		{
			name: "invalid metadata length",
			msg: types.MsgCreateCampaign{
				Coordinator:  sample.Address(r),
				CampaignName: sample.CampaignName(r),
				TotalSupply:  sample.TotalSupply(r),
				Metadata:     sample.Metadata(r, spntypes.MaxMetadataLength+1),
			},
			err: types.ErrInvalidMetadataLength,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
