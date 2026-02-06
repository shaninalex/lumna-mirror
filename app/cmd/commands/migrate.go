package commands

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/models"
	"go.uber.org/dig"
	"gorm.io/gorm"
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

			if err := c.Invoke(migrate); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func migrate(db *gorm.DB, config *config.Config) {
	if err := db.AutoMigrate(
		&models.Identity{},
		&models.Credential{},
		&models.RefreshToken{},

		&models.Project{},
		&models.Board{},
		&models.Column{},
		&models.Task{},
	); err != nil {
		log.Fatal(err)
	}

	log.Printf("Database migrated in %s\n", config.Database.Url)
}
