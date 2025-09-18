// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	orgApp "gitlab.com/shaninalex/flowreon/apps/org/api"
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
	router := web.NewAppRouter(sqlDB, "org")
	kratosClient := kratos.NewKratosService(config.String("kratos.url_browser"))
	router.Use(web.NewAuthMiddleware(kratosClient).Wrap)
	orgApp.NewOrganizationController(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", config.String("org.port")),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server error: %v\n", err)
	}
}
