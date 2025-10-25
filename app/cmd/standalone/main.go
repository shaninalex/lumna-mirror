// Copyright © 2025 Lumna. All rights reserved.

package main

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"net/http"

	authApp "github.com/shaninalex/lumna/app/apps/auth"
	projectApp "github.com/shaninalex/lumna/app/apps/project/api"
	taskApp "github.com/shaninalex/lumna/app/apps/task/api"
	userApp "github.com/shaninalex/lumna/app/apps/user/api"
	"github.com/shaninalex/lumna/app/internal/base"
	"github.com/shaninalex/lumna/app/internal/db"
	"github.com/shaninalex/lumna/app/internal/dir"
	"github.com/shaninalex/lumna/app/internal/web"
	"github.com/shaninalex/lumna/app/startup"
)

//go:embed all:web/browser
var webFS embed.FS

func main() {
	if err := dir.MakeProjectDirectories(); err != nil {
		panic(err)
	}

	config := base.GetConfig()

	sqlDB, err := sql.Open("sqlite3", config.String("database_path"))
	if err != nil {
		panic(err)
	}
	db.ApplyMigrationsEmbed(sqlDB)
	static, _ := fs.Sub(webFS, "web/browser")

	if startup.IsNew(sqlDB) {
		initializer := startup.NewStartup(sqlDB)
		if err = initializer.Run(); err != nil {
			panic(err)
		}
	}

	router := web.DefaultRouter(sqlDB)

	// Public controllers
	router.GET("/", frontendHandler(static))
	authApp.NewAuthController(router)

	// Private controllers
	router.Use(web.NewTokenMiddleware().Wrap)
	userApp.NewUserController(router)
	projectApp.NewProjectController(router)
	taskApp.NewTaskController(router)

	if err = router.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v\n", err))
	}
}
