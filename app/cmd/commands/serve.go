package commands

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/internal/client"
	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/internal/persistence"
	"gitlab.com/shaninalex/lumna/app/web"
	"go.uber.org/dig"
)

func NewRootServeCommand() (cmd *cobra.Command) {
	cmd = &cobra.Command{
		Use:   "serve",
		Short: "Run webserver",
		Args:  cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			c := dig.New()
			configPath, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}
			_ = c.Provide(config.ProvideConfig(configPath))
			_ = c.Provide(persistence.ProvideDB)
			_ = c.Provide(web.NewWebApplication)

			router := web.NewWebApplication(client)
			srv := &http.Server{
				Addr:    fmt.Sprintf(":%d", client.Config().Serve.Port),
				Handler: router,
			}

			log.Printf("Run server on :%d\n", client.Config().Serve.Port)
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("listen: %s\n", err)
				}
			}()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
			<-quit
			log.Println("Shutting down server...")

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			if err := srv.Shutdown(ctx); err != nil {
				log.Fatal("Server forced to shutdown:", err)
			}

			log.Println("Server exiting")

		},
	}

	return cmd
}
