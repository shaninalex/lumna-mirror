package main

import (
	"fmt"
	"os"

	orgApp "gitlab.com/shaninalex/jajirra/apps/org/app"
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
	router := web.DefaultRouter(db, "org")
	kratosClient := kratos.NewKratosService(config.String("kratos.url_browser"))
	router.Use(web.NewAuthMiddleware(kratosClient))
	orgApp.NewOrganizationController(router)
	if err := router.Listen(fmt.Sprintf(":%s", config.String("org.port"))); err != nil {
		panic(err)
	}
}
