package identities

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
)

func NewIdentitiesCreateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create identities",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
			if err != nil {
				panic(err)
			}
			c.DBConnect()

			email := args[0]
			password := args[1]

			user, err := c.UserManager.CreateUser(c.Context(), email, password)
			if err != nil {
				fmt.Println(err)
				return
			}

			log.Println("New user created")
			log.Println("ID: ", user.GetId())
			log.Println("Email: ", user.GetEmail())
		},
	}

	return cmd
}
