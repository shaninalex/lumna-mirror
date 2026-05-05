package commands

import (
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/pkg/persistence"
	"go.uber.org/dig"
)

func NewMigrateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Apply migrations",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			if err := c.Invoke(persistence.ApplyMigrations); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}
