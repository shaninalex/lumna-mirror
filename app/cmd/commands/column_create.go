package commands

import (
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewColumnCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [board_id] [title]",
		Short: "Create column",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClientForCLI(cmd)
			if err != nil {
				panic(err)
			}
			db := c.DB()
			boardID := uuid.MustParse(args[0])
			title := args[1]

			column := models.Column{
				Title:   title,
				BoardID: boardID,
			}
			if result := db.Create(&column); result.Error != nil {
				panic(result.Error)
			}

			log.Println("Column created with ID:", column.ID.String())
		},
	}

	return cmd
}
