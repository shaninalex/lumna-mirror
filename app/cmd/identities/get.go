package identities

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
)

func NewIdentitiesGetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [id]",
		Short: "Get identity by id",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
			if err != nil {
				panic(err)
			}
			c.DBConnect()

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				panic(err)
			}

			user, err := c.UserManager.GetUser(c.Context(), uint(id))
			if err != nil {
				panic(err)
			}

			log.Println("ID: ", user.GetId())
			log.Println("Email: ", user.GetEmail())
		},
	}

	return cmd
}
