package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewProjectsListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List projects",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClientForCLI(cmd)
			if err != nil {
				log.Fatal(err)
			}
			db := c.DB()

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
		},
	}

	return cmd
}
