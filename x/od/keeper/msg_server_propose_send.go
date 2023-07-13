package keeper

import (
	"context"

	"opendao/x/od/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ProposeSend(goCtx context.Context, msg *types.MsgProposeSend) (*types.MsgProposeSendResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// 申请人需要质押1个open
	applicantAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, applicantAddr, types.ModuleName, sdk.NewCoins(sdk.NewCoin("open", sdk.NewInt(types.ApplicantStake))))
	if err != nil {
		return &types.MsgProposeSendResponse{}, types.ErrInsufficientStake
	}

	// 创建提案并存入区块链
	proposalid := k.AppendProposal(ctx, types.Proposal{
		Title:       msg.Title,
		Description: msg.Description,
		Coins:       msg.Coins,
		Expiry:      ctx.BlockHeight() + types.ExpiryBlocks, // 到期区块
		Applicant:   msg.Creator,
		Beneficiary: msg.Beneficiary,
		Agreed:      []string{},
	})

	return &types.MsgProposeSendResponse{Proposalid: proposalid}, nil
}
