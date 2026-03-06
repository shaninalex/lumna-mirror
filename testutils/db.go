package testutils

import (
	"gitlab.com/shaninalex/lumna/app/pkg/persistence"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ProvideTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared&mode=rwc"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if res := db.Exec("PRAGMA foreign_keys = ON"); res.Error != nil {
		panic(res.Error.Error())
	}
	return db
}

func Migrate(db *gorm.DB) error {
	return persistence.Migrate(db)
}
