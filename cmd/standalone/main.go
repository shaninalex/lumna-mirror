// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package main

import (
	"embed"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	authApp "gitlab.com/shaninalex/flowreon/apps/auth/api"
	orgApp "gitlab.com/shaninalex/flowreon/apps/org/api"
	projectApp "gitlab.com/shaninalex/flowreon/apps/project/api"
	userApp "gitlab.com/shaninalex/flowreon/apps/user/api"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

//go:embed all:web/browser
var webFS embed.FS

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

	static, _ := fs.Sub(webFS, "web/browser")

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	router := web.DefaultRouter(sqlDB, "auth")
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path == "/" {
			path = "index.html"
		}
		_, err = static.Open(strings.TrimPrefix(path, "/"))
		if err == nil {
			http.FileServer(http.FS(static)).ServeHTTP(w, r)
			return
		}
		http.ServeFileFS(w, r, static, "index.html")
	})

	kratosClient := kratos.NewKratosService(config.String("kratos.url_browser"))
	// public routes
	authApp.NewAuthController(config, router, kratosClient)

	router.Use(web.NewAuthMiddleware(kratosClient).Wrap)
	// private routes
	orgApp.NewOrganizationController(router)
	projectApp.NewProjectController(router)
	userApp.NewUserController(router)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":4200",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("server started...")
	if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server error: %v\n", err)
	}
}
