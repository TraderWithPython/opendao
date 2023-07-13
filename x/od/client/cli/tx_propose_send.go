package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"opendao/x/od/types"
)

var _ = strconv.Itoa(0)

func CmdProposeSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "propose-send [title] [description] [beneficiary] [coins]",
		Short: "Broadcast message propose_send",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTitle := args[0]
			argDescription := args[1]
			argBeneficiary := args[2]
			argCoins, err := sdk.ParseCoinsNormalized(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgProposeSend(
				clientCtx.GetFromAddress().String(),
				argTitle,
				argDescription,
				argBeneficiary,
				argCoins,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
