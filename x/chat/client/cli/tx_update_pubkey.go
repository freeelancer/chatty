package cli

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"chatty/x/chat/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdatePubkey() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-pubkey",
		Short: "Broadcast message update-pubkey [pubkey filepath]",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			pubkeyFilePath := args[0]
			if pubkeyFilePath == "" {
				return fmt.Errorf("pubkey cannot be empty")
			}
			pemData, err := GetPemDataBs(pubkeyFilePath)
			if err != nil {
				return err
			}
			pemDataHexStr := hex.EncodeToString(pemData)

			msg := types.NewMsgUpdatePubkey(
				clientCtx.GetFromAddress().String(),
				pemDataHexStr,
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
