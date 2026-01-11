package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app2/cmd/identities"
	"gitlab.com/shaninalex/lumna/app2/cmd/migrate"
	"gitlab.com/shaninalex/lumna/app2/cmd/serve"
)

func NewRootCmd() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use: "lumna",
	}

	cmd.AddCommand(serve.NewRootServeCommand())
	cmd.AddCommand(migrate.NewMigrateRootCmd())
	cmd.AddCommand(identities.NewIdentitiesRootCmd())

	cmd.PersistentFlags().String("config", "", "Configuration path. Required.")
	cmd.MarkPersistentFlagRequired("config")
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
