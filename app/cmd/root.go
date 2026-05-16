package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/commands"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "lumna",
	}

	cmd.AddCommand(commands.NewRootServeCommand())
	cmd.AddCommand(commands.NewMigrateRootCmd())
	cmd.AddCommand(commands.NewIdentitiesRootCmd())
	cmd.AddCommand(commands.NewProjectsRootCmd())
	cmd.AddCommand(commands.NewBoardsRootCmd())
	cmd.AddCommand(commands.NewStatusRootCmd())
	cmd.AddCommand(commands.NewImportRootCmd())
	cmd.AddCommand(commands.NewExportRootCmd())

	cmd.PersistentFlags().String("config", "", "Configuration path. Required.")
	_ = cmd.MarkPersistentFlagRequired("config")
	return cmd
}

// Execute run application
func Execute() int {
	c := NewRootCmd()

	if err := c.Execute(); err != nil {
		fmt.Println(err)
		return 1
	}
	return 0
}
