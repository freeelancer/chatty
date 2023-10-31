package cli

import (
	"fmt"
	"strconv"

	"chatty/x/chat/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateChatMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chat-message [receiver] [message] [encrypted (1 or 0)]",
		Short: "Broadcast message create-chat-message",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			receiver := args[0]
			if receiver == "" {
				return fmt.Errorf("receiver cannot be empty")
			}

			message := args[1]
			if message == "" {
				return fmt.Errorf("message cannot be empty")
			}

			encryptedInt, err := strconv.ParseBool(args[2])
			if err != nil {
				return fmt.Errorf("encrypted must be 1 or 0")
			}

			encrypted := false
			if encryptedInt {
				encrypted = true
			}

			msg := types.NewMsgCreateChatMessage(
				clientCtx.GetFromAddress().String(),
				receiver,
				message,
				encrypted,
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
