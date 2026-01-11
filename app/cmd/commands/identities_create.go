package commands

import (
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

func NewIdentitiesCreateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [email] [fullname]",
		Short: "Create identities",
		Long:  "Require 2 arguments - first: email, second: password. For example:\nlumna identities create test@test.com password",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
			if err != nil {
				log.Fatal(err)
			}
			db := c.DB()

			email := args[0]
			fullname := args[1]
			user := models.Identity{
				ID:       uuid.New(),
				FullName: fullname,
				Email:    email,
				Active:   false,
			}

			result := db.Create(&user)
			if result.Error != nil {
				log.Fatal(result.Error)
			}

			credential := models.Credential{
				IdentityID:     user.ID,
				Provider:       "google",
				ProviderUserID: utils.Pointer[string]("google"),
				Email:          &user.Email,
				PasswordHash:   utils.Pointer[string](""),
			}
			result = db.Create(&credential)
			if result.Error != nil {
				log.Fatal(result.Error)
			}

			log.Println("User created with ID:", user.ID.String())
			log.Println("Credentials created", credential.ID.String())
		},
	}

	return cmd
}
