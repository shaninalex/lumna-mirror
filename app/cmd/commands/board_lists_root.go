package commands

import "github.com/spf13/cobra"

func NewBoardListsRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lists",
		Short: "Manage board lists",
	}

	cmd.AddCommand(NewBoardListCreateCmd())

	return cmd
}
