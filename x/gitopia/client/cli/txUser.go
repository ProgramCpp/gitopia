package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/gitopia/gitopia/x/gitopia/types"
)

func CmdCreateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-user [username]",
		Short: "Create a new user",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsUsername, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUser(clientCtx.GetFromAddress().String(), string(argsUsername))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-user [name] [usernameGithub] [avatarUrl] [email] [bio]",
		Short: "Update a user",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsName := string(args[0])
			argsUsernameGithub := string(args[1])
			argsAvatarUrl := string(args[2])
			argsEmail := string(args[3])
			argsBio := string(args[4])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateUser(clientCtx.GetFromAddress().String(), string(argsName), string(argsUsernameGithub), string(argsAvatarUrl), string(argsEmail), string(argsBio))
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-user [id] [username] [usernameGithub] [avatarUrl] [followers] [following] [repositories] [repositories_archived] [organizations] [starred_repos] [subscriptions] [email] [bio] [createdAt] [updatedAt] [extensions]",
		Short: "Delete a user by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteUser(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdTransferUser() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-user [address]",
		Short: "Transfer the user to new address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			address := string(args[0])

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgTransferUser(clientCtx.GetFromAddress().String(), address)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}