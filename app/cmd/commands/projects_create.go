package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gitlab.com/shaninalex/lumna/app/pkg/persistence"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func NewProjectsCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [title]",
		Short: "Create project",
		Long:  "Require 1 argument - project title. For example:\nlumna projects create test_project_name",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			title := args[0]

			if err := c.Invoke(projectsCreate(title)); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func projectsCreate(title string) func(db *gorm.DB) {
	return func(db *gorm.DB) {
		project := &models.Project{
			Title: title,
		}

		if result := db.Create(&project); result.Error != nil {
			panic(result.Error)
		}

		fmt.Println("Project created:")
		fmt.Println(project.String())
	}
}
