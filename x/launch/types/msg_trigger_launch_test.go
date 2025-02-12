package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgTriggerLaunch_ValidateBasic(t *testing.T) {
	addr := sample.Address(r)
	launchID := uint64(0)

	for _, tc := range []struct {
		desc  string
		msg   types.MsgTriggerLaunch
		valid bool
	}{
		{
			desc:  "should validate valid message",
			msg:   *types.NewMsgTriggerLaunch(addr, launchID, 1000),
			valid: true,
		},
		{
			desc:  "should prevent validate message with invalid coordinator address",
			msg:   *types.NewMsgTriggerLaunch("invalid", launchID, 1000),
			valid: false,
		},
		{
			desc:  "should prevent validate message with remaining time less than 0",
			msg:   *types.NewMsgTriggerLaunch(addr, launchID, 0),
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.msg.ValidateBasic()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
