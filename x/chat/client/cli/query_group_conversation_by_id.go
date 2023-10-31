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

func CmdGroupConversationById() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group-conversation-by-id [id]",
		Short: "Query group-conversation-by-id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			groupId := args[0]
			if groupId == "" {
				return fmt.Errorf("groupId cannot be empty")
			}
			groupIdInt, err := strconv.ParseInt(groupId, 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGroupConversationByIdRequest{
				Id: groupIdInt,
			}

			res, err := queryClient.GroupConversationById(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
