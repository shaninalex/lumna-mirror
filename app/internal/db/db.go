package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(connectionString string) *gorm.DB {
	// NOTE: there are can be different databases - planing also use postgres.
	db, err := gorm.Open(sqlite.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if res := db.Raw("PRAGMA foreign_keys = ON"); res.Error != nil {
		panic(res.Error.Error())
	}

	return db
}
