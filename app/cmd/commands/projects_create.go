package commands

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewProjectsCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [owner] [title]",
		Short: "Create project",
		Long:  "Require 2 arguments - first: owner id, second: project title. For example:\nlumna projects create <owner UUID> test_project_name",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
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
			project := &models.Project{
				Title:   title,
				OwnerID: ownerID,
			}

			if result := db.Create(&project); result.Error != nil {
				panic(result.Error)
			}

			fmt.Println("Project created:")
			fmt.Println(project.String())
		},
	}

	return cmd
}
