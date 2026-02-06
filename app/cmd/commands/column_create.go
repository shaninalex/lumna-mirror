package commands

import (
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/models"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func NewColumnCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [board_id] [title]",
		Short: "Create column",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}
			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			boardID := uuid.MustParse(args[0])
			title := args[1]
			column := models.Column{
				Title:   title,
				BoardID: boardID,
			}
			if err := c.Invoke(createColumn(column)); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func createColumn(column models.Column) func(db *gorm.DB) {
	return func(db *gorm.DB) {

		if result := db.Create(&column); result.Error != nil {
			panic(result.Error)
		}

		log.Println("Column created with ID:", column.ID.String())
	}
}
