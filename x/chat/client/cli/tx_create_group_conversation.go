package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"chatty/x/chat/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateGroupConversation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-group-conversation [name] [participants separated by space] [message] [pubkeyFilePath]",
		Short: "Broadcast message create-group-conversation",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			name := args[0]
			if name == "" {
				return fmt.Errorf("name cannot be empty")
			}
			participantsList := args[1]
			if participantsList == "" {
				return fmt.Errorf("participantsList cannot be empty")
			}
			participants := strings.Split(participantsList, " ")

			message := args[2]
			pubkeyFilePath := args[3]
			encrypted := false
			pubkey := ""
			if pubkeyFilePath != "" {
				encrypted = true
				pubkeyBs, err := GetPemDataBs(pubkeyFilePath)
				if err != nil {
					return err
				}
				if message != "" {
					message, err = EncryptMessageWithPubKey(message, pubkeyFilePath)
					if err != nil {
						return err
					}
				}
				pubkey = hex.EncodeToString(pubkeyBs)
			}

			msg := types.NewMsgCreateGroupConversation(
				clientCtx.GetFromAddress().String(),
				name,
				message,
				pubkey,
				participants,
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
