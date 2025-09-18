// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package main

import (
	"fmt"
	"os"

	authApp "gitlab.com/shaninalex/flowreon/apps/auth/api"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
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
	db := database.InitDB(database.BuildDSN(config))

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	router := web.NewAppRouter(sqlDB, "auth")
	kratosClient := kratos.NewKratosService(config.String("kratos.url_browser"))
	authApp.NewAuthController(config, router, kratosClient)

	if err := router.Listen(fmt.Sprintf(":%s", config.String("auth.port"))); err != nil {
		panic(err)
	}
}
