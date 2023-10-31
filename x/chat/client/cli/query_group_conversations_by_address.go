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

func CmdGroupConversationsByAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group-conversations-by-address [address]",
		Short: "Query group-conversations-by-address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			address := args[0]
			if address == "" {
				return fmt.Errorf("address cannot be empty")
			}

			params := &types.QueryGroupConversationsByAddressRequest{Address: address}

			res, err := queryClient.GroupConversationsByAddress(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
