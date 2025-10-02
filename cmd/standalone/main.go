// Copyright © 2025 Lumna. All rights reserved.

package main

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"net/http"

	authApp "gitlab.com/shaninalex/flowreon/apps/auth"
	projectApp "gitlab.com/shaninalex/flowreon/apps/project/api"
	userApp "gitlab.com/shaninalex/flowreon/apps/user/api"
	"gitlab.com/shaninalex/flowreon/internal/db"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

//go:embed all:web/browser
var webFS embed.FS

func main() {
	sqlDB, err := sql.Open("sqlite3", "file:flowreon.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	db.ApplyMigrationsEmbed(sqlDB)
	static, _ := fs.Sub(webFS, "web/browser")

	router := web.DefaultRouter(sqlDB)
	router.GET("/", frontendHandler(static))

	authApp.NewAuthController(router)

	router.Use(web.NewTokenMiddleware().Wrap)
	userApp.NewUserController(router)
	projectApp.NewProjectController(router)
	// other private apps.

	if err = router.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v\n", err))
	}
}
