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

func CmdCreateGroupConversationMessage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-group-conversation-message [group_id] [message] [pubkey filepath]",
		Short: "Broadcast message create-group-conversation-message",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			groupId := args[0]
			if groupId == "" {
				return fmt.Errorf("groupId cannot be empty")
			}
			groupIdInt, err := strconv.ParseInt(groupId, 10, 64)
			if err != nil {
				return err
			}

			message := args[1]
			if message == "" {
				return fmt.Errorf("message cannot be empty")
			}
			pubkeyFilePath := args[2]
			encrypted := false
			if pubkeyFilePath != "" {
				encrypted = true
				if message != "" {
					message, err = EncryptMessageWithPubKey(message, pubkeyFilePath)
					if err != nil {
						return err
					}
				}
			}

			msg := types.NewMsgCreateGroupConversationMessage(
				clientCtx.GetFromAddress().String(),
				message,
				groupIdInt,
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
