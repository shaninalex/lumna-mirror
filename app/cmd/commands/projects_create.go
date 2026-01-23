package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewProjectsCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [title]",
		Short: "Create project",
		Long:  "Require 2 arguments - first: owner id, second: project title. For example:\nlumna projects create <owner UUID> test_project_name",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClientForCLI(cmd)
			if err != nil {
				log.Fatal(err)
			}
			db := c.DB()
			title := args[0]

			project := &models.Project{
				Title: title,
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
