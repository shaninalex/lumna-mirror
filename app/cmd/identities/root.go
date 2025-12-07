package identities

import (
	"github.com/spf13/cobra"
)

func NewIdentitiesRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "identities",
		Short: "Manage identities",
	}

	cmd.AddCommand(NewIdentitiesCreateRootCmd())
	cmd.AddCommand(NewIdentitiesGetRootCmd())
	cmd.AddCommand(NewIdentitiesInviteRootCmd())

	return cmd
}
