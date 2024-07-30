package cli

import (
	"github.com/CudoVentures/cudos-node/x/mail/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

func CmdSend() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [to] [subject] [body]",
		Short: "Send a new mail",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argTo := args[0]
			argSubject := args[1]
			argBody := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSend(clientCtx.GetFromAddress().String(), argTo, argSubject, argBody)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
