package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	authApp "gitlab.com/shaninalex/lumna/_old_app/apps/auth"
	commentApp "gitlab.com/shaninalex/lumna/_old_app/apps/comment/api"
	projectApp "gitlab.com/shaninalex/lumna/_old_app/apps/project/api"
	taskApp "gitlab.com/shaninalex/lumna/_old_app/apps/task/api"
	userApp "gitlab.com/shaninalex/lumna/_old_app/apps/user/api"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/base"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/db"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/web"
	"gitlab.com/shaninalex/lumna/_old_app/startup"
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
	commentApp.NewCommentController(router)

	if base.IsDebug() {
		log.Printf("Configuration path: %s", config.GetConfigPath())
		log.Printf("Database path: %s", dbPath)
	}

	port := config.Int("port")
	if err = router.Run(port); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(fmt.Errorf("server error: %v\n", err))
	}
}
