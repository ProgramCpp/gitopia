package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/gitopia/gitopia/x/gitopia/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdListOrganization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-organization",
		Short: "list all organization",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllOrganizationRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.OrganizationAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowOrganization() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-organization [id]",
		Short: "shows a organization",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			params := &types.QueryGetOrganizationRequest{
				Id: id,
			}

			res, err := queryClient.Organization(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}