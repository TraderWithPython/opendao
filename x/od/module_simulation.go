package od

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"opendao/testutil/sample"
	odsimulation "opendao/x/od/simulation"
	"opendao/x/od/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = odsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgProposeSend = "op_weight_msg_propose_send"
	// TODO: Determine the simulation weight value
	defaultWeightMsgProposeSend int = 100

	opWeightMsgVote = "op_weight_msg_vote"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVote int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	odGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&odGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgProposeSend int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgProposeSend, &weightMsgProposeSend, nil,
		func(_ *rand.Rand) {
			weightMsgProposeSend = defaultWeightMsgProposeSend
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgProposeSend,
		odsimulation.SimulateMsgProposeSend(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVote int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVote, &weightMsgVote, nil,
		func(_ *rand.Rand) {
			weightMsgVote = defaultWeightMsgVote
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVote,
		odsimulation.SimulateMsgVote(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
