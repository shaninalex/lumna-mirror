package commands

import (
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
	"gitlab.com/shaninalex/lumna/app/models"
)

func NewIdentitiesDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete identity",
		Long:  "Require 1 arguments - Identity Id to delete. For example:\nlumna identities delete <uuid>",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
			if err != nil {
				log.Fatal(err)
			}
			id := uuid.MustParse(args[0])

			if result := c.DB().Where("id = ?", id.String()).Delete(&models.Identity{}); result.Error != nil {
				log.Fatal(result.Error)
			}

			log.Println("User with id", id.String(), "deleted.")
		},
	}

	return cmd
}
