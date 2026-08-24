package commands

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func NewBoardsCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [project] [title]",
		Short: "Create board",
		Long:  "Require 2 arguments - first: project id, second: board title. For example:\nlumna boards create <project UUID> test_board_name",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}
			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			strId := args[0]
			title := args[1]

			ownerID, err := strconv.Atoi(strId)
			if err != nil {
				panic(err)
			}
			board := models.Board{
				Title:     title,
				ProjectID: uint(ownerID),
			}

			if err := c.Invoke(createBoard(board)); err != nil {
				panic(board)
			}

		},
	}

	return cmd
}

func createBoard(board models.Board) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		if result := db.Create(&board); result.Error != nil {
			panic(result.Error)
		}

		fmt.Println("Board created:")
		fmt.Println(board.String())
	}
}
