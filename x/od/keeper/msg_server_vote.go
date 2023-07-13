package keeper

import (
	"context"

	"opendao/x/od/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/duke-git/lancet/v2/slice"
)

func (k msgServer) Vote(goCtx context.Context, msg *types.MsgVote) (*types.MsgVoteResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 判断是否是MP
	if !slice.Contain(types.MPs, msg.Creator) {
		return &types.MsgVoteResponse{}, types.ErrNotMP
	}

	// 添加同意
	proposal, found := k.GetProposal(ctx, msg.Proposalid)
	if !found {
		return &types.MsgVoteResponse{}, types.ErrProposalNotFound
	}
	proposal.Agreed = append(proposal.Agreed, msg.Creator)
	k.SetProposal(ctx, proposal)

	// 判断是否过半
	threshold := (len(types.MPs) + 2 - 1) / 2 // 总数除2向上取整
	if len(proposal.Agreed) >= threshold {
		// 归还提案人质押的open
		applicantAddr, _ := sdk.AccAddressFromBech32(proposal.Applicant)
		err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, applicantAddr, sdk.NewCoins(sdk.NewCoin("open", sdk.NewInt(types.ApplicantStake))))
		if err != nil {
			return &types.MsgVoteResponse{}, err
		}
		// 将mint coins给beneficiary
		err = k.bankKeeper.MintCoins(ctx, types.ModuleName, proposal.Coins)
		if err != nil {
			return &types.MsgVoteResponse{}, err
		}
		beneficiaryAddr, err := sdk.AccAddressFromBech32(proposal.Beneficiary)
		if err != nil {
			return &types.MsgVoteResponse{}, err
		}
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, beneficiaryAddr, proposal.Coins)
		if err != nil {
			return &types.MsgVoteResponse{}, err
		}
	}

	return &types.MsgVoteResponse{}, nil
}
