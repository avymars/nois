package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateCoordinatorDescription(
	goCtx context.Context,
	msg *types.MsgUpdateCoordinatorDescription,
) (*types.MsgUpdateCoordinatorDescriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the coordinator address is already in the store
	coordByAddress, found := k.getCoordinatorByAddress(ctx, msg.Address)
	if !found {
		return &types.MsgUpdateCoordinatorDescriptionResponse{},
			sdkerrors.Wrap(types.ErrCoordAddressNotFound, msg.Address)
	}

	coord, found := k.GetCoordinator(ctx, coordByAddress.CoordinatorID)
	if !found {
		return &types.MsgUpdateCoordinatorDescriptionResponse{},
			spnerrors.Criticalf("a coordinator address is associated to a non-existent coordinator ID: %d",
				coordByAddress.CoordinatorID)
	}

	// Check if the coordinator is inactive
	if !coord.Active {
		return &types.MsgUpdateCoordinatorDescriptionResponse{},
			spnerrors.Criticalf("inactive coordinator address should not exist in store, ID: %d",
				coordByAddress.CoordinatorID)
	}

	if len(msg.Description.Identity) > 0 {
		coord.Description.Identity = msg.Description.Identity
	}
	if len(msg.Description.Website) > 0 {
		coord.Description.Website = msg.Description.Website
	}
	if len(msg.Description.Details) > 0 {
		coord.Description.Details = msg.Description.Details
	}

	k.SetCoordinator(ctx, coord)

	return &types.MsgUpdateCoordinatorDescriptionResponse{}, nil
}
