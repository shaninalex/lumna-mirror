package main

import (
	"fmt"
	"os"

	projectApp "gitlab.com/shaninalex/flowreon/apps/project/app"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/web"
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
	router := web.AuthRouter(config, db, "org")
	projectApp.NewProjectController(router)
	if err := router.Listen(fmt.Sprintf(":%s", config.String("project.port"))); err != nil {
		panic(err)
	}
}
