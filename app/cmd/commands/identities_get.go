package commands

import (
	"log"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func NewIdentitiesGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [id]",
		Short: "Get identity by id",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id := uuid.MustParse(args[0])
			log.Println("Get identity by id:", id.String())
		},
	}

	return cmd
}
