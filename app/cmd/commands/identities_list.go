package commands

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewIdentitiesListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Identities list",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClientForCLI(cmd)
			if err != nil {
				log.Fatal(err)
			}
			db := c.DB()
			identities := []models.Identity{}
			if result := db.Find(&identities); result.Error != nil {
				panic(result.Error)
			}

			for i, identity := range identities {
				fmt.Println(i+1, identity.String())
			}
		},
	}

	return cmd
}
