package commands

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewBoardsCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [project] [title]",
		Short: "Create board",
		Long:  "Require 2 arguments - first: project id, second: board title. For example:\nlumna boards create <project UUID> test_board_name",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClientForCLI(cmd)
			if err != nil {
				log.Fatal(err)
			}
			db := c.DB()
			strId := args[0]
			title := args[1]

			ownerID, err := uuid.Parse(strId)
			if err != nil {
				panic(err)
			}
			board := &models.Board{
				Title:     title,
				ProjectID: ownerID,
			}

			if result := db.Create(&board); result.Error != nil {
				panic(result.Error)
			}

			fmt.Println("Board created:")
			fmt.Println(board.String())
		},
	}

	return cmd
}
