package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/models"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func NewProjectsListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}

			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			if err := c.Invoke(projectsList); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func projectsList(db *gorm.DB) {
	projects := []models.Project{}

	if result := db.Preload("Boards").Find(&projects); result.Error != nil {
		panic(result.Error)
	}

	for i, p := range projects {
		fmt.Println(i, p.String())
		for _, b := range p.Boards {
			fmt.Println("\t\t", b.String())
		}
		fmt.Println()
	}

	fmt.Println("Total projects: ", len(projects))
}
