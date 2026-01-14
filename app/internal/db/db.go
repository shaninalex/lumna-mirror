package db

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/internal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(connectionString string) *gorm.DB {
	// NOTE: there are can be different databases - planing also use postgres. But not for now.
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if res := db.Exec("PRAGMA foreign_keys = ON"); res.Error != nil {
		panic(res.Error.Error())
	}

	return db
}

func GetDB(ctx context.Context) *gorm.DB {
	conn, ok := ctx.Value(internal.ContextDB).(*gorm.DB)
	if !ok {
		panic("db not found in context")
	}
	return conn
}
