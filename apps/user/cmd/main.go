// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package main

import (
	"fmt"
	"os"

	userController "gitlab.com/shaninalex/flowreon/apps/user/api"
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
	db := database.InitDB(database.BuildDSN(config))
	router := web.AuthRouter(config, db, "user")
	userController.NewUserController(router)
	if err := router.Listen(fmt.Sprintf(":%s", config.String("user.port"))); err != nil {
		panic(err)
	}
}
