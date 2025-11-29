package tests

import (
	"context"
	"database/sql"
	"sync"

	_ "github.com/mattn/go-sqlite3"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

var (
	once  sync.Once
	sqlDB *sql.DB
)

func SharedDatabase() *sql.DB {
	once.Do(func() {
		var err error
		sqlDB, err = sql.Open("sqlite3", ":memory:")
		if err != nil {
			panic(err)
		}
		db.ApplyMigrationsEmbed(sqlDB)
	})

	return sqlDB

}

func TestContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, db.CONTEXT_DB, SharedDatabase())

	return ctx
}

func ResetDatabase() {
	tables := []string{"users"}
	for _, t := range tables {
		sqlDB.Exec("DELETE FROM " + t)
	}
}
