package commands

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewMigrateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Apply migrations",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
			if err != nil {
				log.Fatal(err)
			}
			db := c.DB()

			if err = db.AutoMigrate(
				&models.Identity{},
				&models.Credential{},
				&models.RefreshToken{},
			); err != nil {
				log.Fatal(err)
			}

			log.Println("Database migrated")
		},
	}

	return cmd
}
