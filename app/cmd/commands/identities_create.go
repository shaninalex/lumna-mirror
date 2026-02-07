package commands

import (
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/auth"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/internal/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

func NewIdentitiesCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [email] [fullname] [password] [active]",
		Short: "Create identities",
		Long:  "Require 2 arguments - first: email, second: password. For example:\nlumna identities create test@test.com password",
		Args:  cobra.MinimumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}
			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)

			email := args[0]
			fullname := args[1]
			pwd := args[2]
			str_active := args[3]

			intActive, err := strconv.Atoi(str_active)
			if err != nil {
				log.Fatal(err)
			}

			active := false
			if intActive == 1 {
				active = true
			}

			if err := c.Invoke(identityCreate(email, fullname, pwd, active)); err != nil {
				panic(err)
			}
		},
	}

	return cmd
}

func identityCreate(email, fullname, pwd string, active bool) func(db *gorm.DB) {
	return func(db *gorm.DB) {

		user := models.Identity{
			ID:       uuid.New(),
			FullName: fullname,
			Email:    email,
			Active:   active,
		}

		if result := db.Create(&user); result.Error != nil {
			log.Fatal(result.Error)
		}

		pwdHash, err := auth.CreatePasswordHash(pwd)
		if err != nil {
			log.Fatal(err)
		}

		credential := models.Credential{
			IdentityID:     user.ID,
			Provider:       "local",
			ProviderUserID: utils.Pointer(user.ID.String()),
			Email:          &user.Email,
			PasswordHash:   utils.Pointer(pwdHash),
		}

		if result := db.Create(&credential); result.Error != nil {
			log.Fatal(result.Error)
		}

		log.Println("User created with ID:", user.ID.String())
		log.Println("Credentials created", credential.ID.String())
	}
}
