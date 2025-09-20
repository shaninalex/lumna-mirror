// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package main

import (
	"database/sql"
	"embed"
	"errors"
	"io/fs"
	"log"
	"net/http"

	"gitlab.com/shaninalex/flowreon/apps/auth"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

//go:embed all:web/browser
var webFS embed.FS

func main() {
	sqlDB, err := sql.Open("sqlite3", "file:flowreon.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	database.ApplyMigrationsEmbed(sqlDB)
	static, _ := fs.Sub(webFS, "web/browser")

	router := web.DefaultRouter(sqlDB, "standalone")
	router.GET("/", frontendHandler(static))

	sessionStore := web.NewCookieStoreDatabase(sqlDB, []byte("very-secret-key"))
	auth.NewAuthController(router, sessionStore)

	log.Println("server started...")
	if err := router.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("Server error: %v\n", err)
	}
}
