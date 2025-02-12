package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) TriggerLaunch(goCtx context.Context, msg *types.MsgTriggerLaunch) (*types.MsgTriggerLaunchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.LaunchID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrChainNotFound, "%d", msg.LaunchID)
	}

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if chain.CoordinatorID != coordID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the chain is %d",
			chain.CoordinatorID,
		))
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrapf(types.ErrTriggeredLaunch, "%d", msg.LaunchID)
	}

	if msg.RemainingTime < k.LaunchTimeRange(ctx).MinLaunchTime {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooLow, "%d", msg.RemainingTime)
	}
	if msg.RemainingTime > k.LaunchTimeRange(ctx).MaxLaunchTime {
		return nil, sdkerrors.Wrapf(types.ErrLaunchTimeTooHigh, "%d", msg.RemainingTime)
	}

	// set launch timestamp
	chain.LaunchTriggered = true
	timestamp := ctx.BlockTime().Unix() + msg.RemainingTime
	chain.LaunchTimestamp = timestamp

	// set revision height for monitoring IBC client
	chain.ConsumerRevisionHeight = ctx.BlockHeight()

	k.SetChain(ctx, chain)

	err = ctx.EventManager().EmitTypedEvent(&types.EventLaunchTriggered{
		LaunchID:        msg.LaunchID,
		LaunchTimestamp: timestamp,
	})

	return &types.MsgTriggerLaunchResponse{}, err
}
