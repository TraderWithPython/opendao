package od_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "opendao/testutil/keeper"
	"opendao/testutil/nullify"
	"opendao/x/od"
	"opendao/x/od/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ProposalList: []types.Proposal{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		ProposalCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OdKeeper(t)
	od.InitGenesis(ctx, *k, genesisState)
	got := od.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ProposalList, got.ProposalList)
	require.Equal(t, genesisState.ProposalCount, got.ProposalCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
