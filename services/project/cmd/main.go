package main

import (
	"fmt"
	"os"

	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/kratos"
	"gitlab.com/shaninalex/jajirra/internal/web"
	taskApp "gitlab.com/shaninalex/jajirra/services/project/app"
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
	router := web.DefaultRouter(db, "task")
	kratosClient := kratos.NewKratosService(config.String("kratos.url_browser"))
	router.Use(web.NewAuthMiddleware(kratosClient))
	taskApp.NewTaskController(router)
	if err := router.Listen(fmt.Sprintf(":%s", config.String("task.port"))); err != nil {
		panic(err)
	}
}
