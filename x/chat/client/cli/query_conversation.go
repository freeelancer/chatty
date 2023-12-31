package cli

import (
	"fmt"
	"strconv"

	"chatty/x/chat/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdConversation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "conversation",
		Short: "Query conversation [chatterA] [chatterB]",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			chatterA := args[0]
			if chatterA == "" {
				return fmt.Errorf("chatterA cannot be empty")
			}

			chatterB := args[1]
			if chatterB == "" {
				return fmt.Errorf("chatterB cannot be empty")
			}

			params := &types.QueryConversationRequest{
				ChatterA: chatterA,
				ChatterB: chatterB,
			}

			res, err := queryClient.Conversation(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
