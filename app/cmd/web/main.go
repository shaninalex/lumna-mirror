package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"gitlab.com/shaninalex/lumna/app/base"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/web"

	authApp "gitlab.com/shaninalex/lumna/app/web/api/auth"
	tokenApp "gitlab.com/shaninalex/lumna/app/web/api/token"
	userApp "gitlab.com/shaninalex/lumna/app/web/api/user"
)

func main() {
	config := base.GetConfig()
	fmt.Println("Run Lumna as a webserver")

	dbConnection, err := sql.Open("sqlite3", config.String("database_path"))
	if err != nil {
		panic(err)
	}

	router := web.NewDefaultRouter()
	router.ApplyMiddlewares([]web.RouterMiddleware{
		db.NewMiddleware(dbConnection),
	})

	static := web.GetStaticFS()

	// Public controllers
	if static != nil {
		router.GET("/", web.FrontendHandler(static))
	}

	authApp.RegisterAuthController(router)   // /api/v1/auth/...
	userApp.RegisterUserController(router)   // /api/v1/user/...
	tokenApp.RegisterTokenController(router) // /api/v1/token/...

	port := config.Int("port")
	if err := router.Run(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v", err))
	}
}
