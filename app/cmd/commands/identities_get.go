package commands

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func NewIdentitiesGetRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [id]",
		Short: "Get identity by id",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				panic(err)
			}

			log.Println("Get identity by id:", id)
		},
	}

	return cmd
}
