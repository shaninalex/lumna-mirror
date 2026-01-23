package commands

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewMigrateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Apply migrations",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClientForCLI(cmd)
			if err != nil {
				log.Fatal(err)
			}
			db := c.DB()

			if err = db.AutoMigrate(
				&models.Identity{},
				&models.Credential{},
				&models.RefreshToken{},
				&models.Project{},
				&models.Board{},
				&models.List{},
				&models.Task{},
			); err != nil {
				log.Fatal(err)
			}

			log.Println("Database migrated")
		},
	}

	return cmd
}
