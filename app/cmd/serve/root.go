package serve

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/cmd/client"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/web"
	"gitlab.com/shaninalex/lumna/app/web/middlewares"

	authApp "gitlab.com/shaninalex/lumna/app/web/api/auth"
	tokenApp "gitlab.com/shaninalex/lumna/app/web/api/token"
	userApp "gitlab.com/shaninalex/lumna/app/web/api/user"
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
	c, err := client.NewClient(cmd)
	if err != nil {
		panic(err)
	}
	conn := c.DBConnect()

	embed, err := cmd.Flags().GetBool("embed")
	if err != nil {
		panic(fmt.Errorf("unable to set embed"))
	}

	router := web.NewDefaultRouter()
	router.ApplyMiddlewares([]web.RouterMiddleware{
		middlewares.NewRecoveryMiddleware(),
		db.NewMiddleware(conn),
		middlewares.NewLoggerMiddleware(),
		middlewares.NewCommonMiddleware(),
		middlewares.NewHeadersMiddleware(),
	})

	if embed {
		if static := web.GetStaticFS(); static != nil {
			router.GET("/", web.FrontendHandler(static))
		}
	}

	authApp.RegisterAuthController(router) // /api/v1/auth/...

	router.Use(middlewares.NewTokenMiddleware().Wrap)
	userApp.RegisterUserController(router)   // /api/v1/user/...
	tokenApp.RegisterTokenController(router) // /api/v1/token/...

	if err := router.Run(c.Config().Serve.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v", err))
	}
}
