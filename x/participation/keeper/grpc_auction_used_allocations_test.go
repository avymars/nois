package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/participation/types"
)

func TestAuctionUsedAllocationsQuerySingle(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(sdkCtx)
	msgs := createNAuctionUsedAllocations(tk.ParticipationKeeper, sdkCtx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAuctionUsedAllocationsRequest
		response *types.QueryGetAuctionUsedAllocationsResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetAuctionUsedAllocationsRequest{
				Address:   msgs[0].Address,
				AuctionID: msgs[0].AuctionID,
			},
			response: &types.QueryGetAuctionUsedAllocationsResponse{AuctionUsedAllocations: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetAuctionUsedAllocationsRequest{
				Address:   msgs[1].Address,
				AuctionID: msgs[1].AuctionID,
			},
			response: &types.QueryGetAuctionUsedAllocationsResponse{AuctionUsedAllocations: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetAuctionUsedAllocationsRequest{
				Address:   strconv.Itoa(100000),
				AuctionID: 100000,
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ParticipationKeeper.AuctionUsedAllocations(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestAuctionUsedAllocationsQueryPaginated(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(sdkCtx)
	msgs := createNAuctionUsedAllocationsWithSameAddress(tk.ParticipationKeeper, sdkCtx, 5)
	address := msgs[0].Address

	request := func(addr string, next []byte, offset, limit uint64, total bool) *types.QueryAllAuctionUsedAllocationsRequest {
		return &types.QueryAllAuctionUsedAllocationsRequest{
			Address: addr,
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("EmptySet", func(t *testing.T) {
		resp, err := tk.ParticipationKeeper.AuctionUsedAllocationsAll(wctx, request("invalid", nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, 0, int(resp.Pagination.Total))
		require.Nil(t, resp.AuctionUsedAllocations)
	})
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ParticipationKeeper.AuctionUsedAllocationsAll(wctx, request(address, nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AuctionUsedAllocations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AuctionUsedAllocations),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := tk.ParticipationKeeper.AuctionUsedAllocationsAll(wctx, request(address, next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.AuctionUsedAllocations), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.AuctionUsedAllocations),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := tk.ParticipationKeeper.AuctionUsedAllocationsAll(wctx, request(address, nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.AuctionUsedAllocations),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := tk.ParticipationKeeper.AuctionUsedAllocationsAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
