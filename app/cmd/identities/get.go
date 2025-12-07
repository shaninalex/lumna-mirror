package identities

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
)

func NewIdentitiesGetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get identity by id",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
			if err != nil {
				panic(err)
			}

			log.Println("Get identity", c)

		},
	}

	return cmd
}
