package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "opendao/testutil/keeper"
	"opendao/testutil/nullify"
	"opendao/x/od/keeper"
	"opendao/x/od/types"
)

func createNProposal(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Proposal {
	items := make([]types.Proposal, n)
	for i := range items {
		items[i].Id = keeper.AppendProposal(ctx, items[i])
	}
	return items
}

func TestProposalGet(t *testing.T) {
	keeper, ctx := keepertest.OdKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetProposal(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestProposalRemove(t *testing.T) {
	keeper, ctx := keepertest.OdKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveProposal(ctx, item.Id)
		_, found := keeper.GetProposal(ctx, item.Id)
		require.False(t, found)
	}
}

func TestProposalGetAll(t *testing.T) {
	keeper, ctx := keepertest.OdKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllProposal(ctx)),
	)
}

func TestProposalCount(t *testing.T) {
	keeper, ctx := keepertest.OdKeeper(t)
	items := createNProposal(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetProposalCount(ctx))
}
