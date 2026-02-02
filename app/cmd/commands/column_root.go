package commands

import "github.com/spf13/cobra"

func NewColumnsRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "columns",
		Short: "Manage columns",
	}

	cmd.AddCommand(NewColumnCreateCmd())

	return cmd
}
