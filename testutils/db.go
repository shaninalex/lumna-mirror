package testutils

import (
	"context"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/services/persistence"
	"gorm.io/gorm"
)

func SetupTest() (context.Context, *gorm.DB) {
	ctx := context.Background()
	db := ProvideTestDB()
	_ = Migrate(db)

	return ctx, db
}

func ProvideTestDB() *gorm.DB {
	return persistence.ProvideDB(ProvideTestConfig())
}

func Migrate(db *gorm.DB) error {
	return persistence.ApplyMigrations(ProvideTestConfig())
}

var tablesToClean = []string{"identities", "credentials"}

func ClearDB(db *gorm.DB) {
	for _, t := range tablesToClean {
		if tx := db.Exec(fmt.Sprintf("DELETE FROM %s", t)); tx.Error != nil {
			panic(tx.Error)
		}
	}
}
