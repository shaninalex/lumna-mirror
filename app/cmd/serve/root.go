package serve

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
	"gitlab.com/shaninalex/lumna/app/base"
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

	cmd.Flags().Int("port", 8000, "port to listen on")
	cmd.Flags().Bool("embed", false, "Embed web client static files")
	return cmd
}

func RunWebServer(cmd *cobra.Command, args []string) {
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		panic(err)
	}

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	_ = ln.Close()

	embed, err := cmd.Flags().GetBool("embed")
	if err != nil {
		panic(fmt.Errorf("unable to set embed"))
	}

	config := base.GetConfig()
	fmt.Println("Run Lumna as a webserver")
	log.Println("Configuration path: ", config.GetConfigPath())
	log.Println("Database path: ", config.String("database_path"))

	dbConnection, err := sql.Open("sqlite3", config.String("database_path"))
	if err != nil {
		panic(err)
	}

	router := web.NewDefaultRouter()
	router.ApplyMiddlewares([]web.RouterMiddleware{
		middlewares.NewRecoveryMiddleware(),
		db.NewMiddleware(dbConnection),
		middlewares.NewLoggerMiddleware(),
		middlewares.NewCommonMiddleware(),
		middlewares.NewHeadersMiddleware(),
	})

	if embed {
		if static := web.GetStaticFS(); static != nil {
			router.GET("/", web.FrontendHandler(static))
		}
	}

	authApp.RegisterAuthController(router)   // /api/v1/auth/...
	userApp.RegisterUserController(router)   // /api/v1/user/...
	tokenApp.RegisterTokenController(router) // /api/v1/token/...

	if err := router.Run(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v", err))
	}
}
