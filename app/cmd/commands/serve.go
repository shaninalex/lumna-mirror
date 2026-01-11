package commands

import (
	"log"

	"github.com/spf13/cobra"
)

func NewRootServeCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "serve",
		Short: "Run webserver",
		Args:  cobra.ArbitraryArgs,
		Run:   RunWebServer,
	}

	cmd.Flags().Bool("embed", false, "Embed web client static files")
	return cmd
}

func RunWebServer(cmd *cobra.Command, args []string) {
	log.Println("Run webserver")
}
