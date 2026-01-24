package commands

import (
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewBoardListCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [board_id] [title]",
		Short: "Create board list",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClientForCLI(cmd)
			if err != nil {
				panic(err)
			}
			db := c.DB()
			boardID := uuid.MustParse(args[0])
			title := args[1]

			boardList := models.BoardList{
				Title:   title,
				BoardID: boardID,
			}
			if result := db.Create(&boardList); result.Error != nil {
				panic(result.Error)
			}

			log.Println("Board list created with ID:", boardList.ID.String())
		},
	}

	return cmd
}
