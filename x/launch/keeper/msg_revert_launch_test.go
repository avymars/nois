package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRevertLaunch(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)

	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.Address(r)
	coordAddress2 := sample.Address(r)
	coordNoExist := sample.Address(r)
	chainIDNoExist := uint64(1000)

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	// Create chains
	msgCreateChain := sample.MsgCreateChain(r, coordAddress, "", false, 0)
	res, err := ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	notLaunched := res.LaunchID

	res, err = ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	delayNotReached := res.LaunchID
	chain, found := tk.LaunchKeeper.GetChain(sdkCtx, delayNotReached)
	require.True(t, found)
	chain.LaunchTriggered = true
	chain.LaunchTimestamp = testkeeper.ExampleTimestamp.Unix() - tk.LaunchKeeper.RevertDelay(sdkCtx) + 1
	tk.LaunchKeeper.SetChain(sdkCtx, chain)

	res, err = ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	delayReached := res.LaunchID
	chain, found = tk.LaunchKeeper.GetChain(sdkCtx, delayReached)
	require.True(t, found)
	chain.LaunchTriggered = true
	chain.LaunchTimestamp = testkeeper.ExampleTimestamp.Unix() - tk.LaunchKeeper.RevertDelay(sdkCtx)
	tk.LaunchKeeper.SetChain(sdkCtx, chain)

	res, err = ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	ibcConnected := res.LaunchID
	chain, found = tk.LaunchKeeper.GetChain(sdkCtx, ibcConnected)
	require.True(t, found)
	chain.LaunchTriggered = true
	chain.LaunchTimestamp = testkeeper.ExampleTimestamp.Unix() - tk.LaunchKeeper.RevertDelay(sdkCtx) + 1
	chain.MonitoringConnected = true
	tk.LaunchKeeper.SetChain(sdkCtx, chain)

	// setup monitoringc keeper with client IDs
	testClientID := "test-client-id-1"
	tk.MonitoringConsumerKeeper.AddVerifiedClientID(sdkCtx, delayReached, testClientID)

	for _, tc := range []struct {
		name     string
		msg      types.MsgRevertLaunch
		clientID string
		err      error
	}{
		{
			name:     "should allow reverting launch if revert delay reached",
			msg:      *types.NewMsgRevertLaunch(coordAddress, delayReached),
			clientID: testClientID,
		},
		{
			name: "should prevent reverting launch if revert delay not reached",
			msg:  *types.NewMsgRevertLaunch(coordAddress, delayNotReached),
			err:  types.ErrRevertDelayNotReached,
		},
		{
			name: "should prevent reverting launch if chain not launched",
			msg:  *types.NewMsgRevertLaunch(coordAddress, notLaunched),
			err:  types.ErrNotTriggeredLaunch,
		},
		{
			name: "should prevent reverting launch with non existent coordinator",
			msg:  *types.NewMsgRevertLaunch(coordNoExist, delayReached),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "should prevent reverting launch with invalid coordinator",
			msg:  *types.NewMsgRevertLaunch(coordAddress2, delayReached),
			err:  profiletypes.ErrCoordInvalid,
		},
		{
			name: "should prevent reverting launch if chain is already ibc connected",
			msg:  *types.NewMsgRevertLaunch(coordAddress, ibcConnected),
			err:  types.ErrChainMonitoringConnected,
		},
		{
			name: "should prevent reverting launch with non existent chain id",
			msg:  *types.NewMsgRevertLaunch(coordAddress, chainIDNoExist),
			err:  types.ErrChainNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Send the message
			_, err := ts.LaunchSrv.RevertLaunch(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// Check value of chain
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, tc.msg.LaunchID)
			require.True(t, found)
			require.False(t, chain.LaunchTriggered)
			require.EqualValues(t, int64(0), chain.LaunchTimestamp)

			// check that monitoringc client ids are removed
			_, found = tk.MonitoringConsumerKeeper.GetVerifiedClientID(sdkCtx, tc.msg.LaunchID)
			require.False(t, found)
			_, found = tk.MonitoringConsumerKeeper.GetLaunchIDFromVerifiedClientID(sdkCtx, tc.clientID)
			require.False(t, found)
		})
	}
}
