package od

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"opendao/x/od/keeper"
	"opendao/x/od/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the proposal
	for _, elem := range genState.ProposalList {
		k.SetProposal(ctx, elem)
	}

	// Set proposal count
	k.SetProposalCount(ctx, genState.ProposalCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ProposalList = k.GetAllProposal(ctx)
	genesis.ProposalCount = k.GetProposalCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
