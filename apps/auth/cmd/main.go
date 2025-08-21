package main

import (
	"fmt"
	"os"

	authApp "gitlab.com/shaninalex/jajirra/apps/auth/app"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/kratos"
	"gitlab.com/shaninalex/jajirra/internal/web"
)

func main() {
	args := os.Args
	var configPath string
	if len(args) < 2 || os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	} else {
		configPath = args[1]
	}

	config := base.NewConfig(configPath)
	db := database.InitDB(config.String("app.dsn"))

	router := web.DefaultRouter(db, "auth")
	kratosClient := kratos.NewKratosService(config.String("kratos.url_browser"))
	NewAuthApi := authApp.NewAuthApi()
	authApp.NewAuthController(config, router, NewAuthApi, kratosClient)

	if err := router.Listen(fmt.Sprintf(":%s", config.String("auth.port"))); err != nil {
		panic(err)
	}
}
