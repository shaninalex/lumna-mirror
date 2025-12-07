package identities

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
)

func NewIdentitiesCreateRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create identities",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
			if err != nil {
				panic(err)
			}

			log.Println("Create identity", c)

		},
	}

	return cmd
}
