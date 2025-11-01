package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	authApp "gitlab.com/shaninalex/lumna/app/apps/auth"
	projectApp "gitlab.com/shaninalex/lumna/app/apps/project/api"
	taskApp "gitlab.com/shaninalex/lumna/app/apps/task/api"
	userApp "gitlab.com/shaninalex/lumna/app/apps/user/api"
	"gitlab.com/shaninalex/lumna/app/internal/base"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/internal/web"
	"gitlab.com/shaninalex/lumna/app/startup"
)

func main() {
	config := base.GetConfig()

	dbPath := db.GetDatabaseUri(config)

	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(err)
	}
	db.ApplyMigrationsEmbed(sqlDB)
	static := getStaticFS()

	if startup.IsNew(sqlDB) {
		initializer := startup.NewStartup(sqlDB)
		if err = initializer.Run(); err != nil {
			panic(err)
		}
	}

	router := web.DefaultRouter(sqlDB)

	// Public controllers
	if static != nil {
		router.GET("/", frontendHandler(static))
	}
	authApp.NewAuthController(router)

	// Private controllers
	router.Use(web.NewTokenMiddleware().Wrap)
	userApp.NewUserController(router)
	projectApp.NewProjectController(router)
	taskApp.NewTaskController(router)

	log.Printf("Configuration path: %s", os.Getenv("LUMNA_CONFIG_PATH"))
	log.Printf("Database path: %s", dbPath)

	port := config.Int("port")
	if err = router.Run(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v\n", err))
	}
}
