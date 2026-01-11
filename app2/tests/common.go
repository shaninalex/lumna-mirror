package tests

import (
	"context"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"

	"gitlab.com/shaninalex/lumna/app2/global"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
	"gitlab.com/shaninalex/lumna/app2/web"
	"gitlab.com/shaninalex/lumna/app2/web/middlewares"
)

var (
	once          sync.Once
	sqlDB         *sql.DB
	TestSecretKey = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func SharedDatabase() *sql.DB {
	once.Do(func() {
		var err error
		sqlDB, err = sql.Open("sqlite3", ":memory:")
		if err != nil {
			panic(err)
		}
		db.ApplyMigrations(sqlDB)

		if _, err = sqlDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
			panic(err)
		}
	})

	return sqlDB

}

func TestContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, global.ContextDB, SharedDatabase())

	return ctx
}

func ResetDatabase() {
	tables := []string{"users", "users_tokens", "projects", "boards", "board_lists"}
	for _, t := range tables {
		sqlDB.Exec("DELETE FROM " + t)
	}
}

func NewTestRouter(ctx context.Context) *web.Router {
	conn := db.FromContext(ctx)
	router := web.NewDefaultRouter()
	router.ApplyMiddlewares([]web.RouterMiddleware{
		db.NewMiddleware(conn),
		middlewares.NewCommonMiddleware(),
		middlewares.NewHeadersMiddleware(),
	})
	return router
}

func Context() context.Context {
	ctx := TestContext()
	ResetDatabase()
	return ctx
}
