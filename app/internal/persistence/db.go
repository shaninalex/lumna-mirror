package persistence

import (
	"log"

	"gitlab.com/shaninalex/lumna/app/internal/config"
	"gitlab.com/shaninalex/lumna/app/models"
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

func ProvideDB(config *config.Config) *gorm.DB {
	return Connect(config.Database.Url)
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Identity{},
		&models.Credential{},
		&models.RefreshToken{},

		&models.Project{},
		&models.Board{},
		&models.Column{},
		&models.Task{},

		&models.Job{},

		&models.ActivityLog{},
	)

	if err != nil {
		return err
	}

	log.Printf("Database migrated")

	return nil
}
