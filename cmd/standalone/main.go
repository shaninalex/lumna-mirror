// Copyright © 2025 Lumna. All rights reserved.

package main

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"net/http"

	authApp "github.com/shaninalex/lumna/apps/auth"
	projectApp "github.com/shaninalex/lumna/apps/project/api"
	taskApp "github.com/shaninalex/lumna/apps/task/api"
	userApp "github.com/shaninalex/lumna/apps/user/api"
	"github.com/shaninalex/lumna/internal/db"
	"github.com/shaninalex/lumna/internal/web"
)

//go:embed all:web/browser
var webFS embed.FS

func main() {
	sqlDB, err := sql.Open("sqlite3", "file:lumna.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	db.ApplyMigrationsEmbed(sqlDB)
	static, _ := fs.Sub(webFS, "web/browser")
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
