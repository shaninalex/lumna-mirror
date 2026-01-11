package commands

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
	"gitlab.com/shaninalex/lumna/app/web"
)

func NewRootServeCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "serve",
		Short: "Run webserver",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			c, err := client.NewClient(cmd)
			if err != nil {
				log.Fatal(err)
			}

			router := web.NewWebApplication(c.Config())

			if err = router.Run(":8000"); err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().Bool("embed", false, "Embed web client static files")
	return cmd
}
