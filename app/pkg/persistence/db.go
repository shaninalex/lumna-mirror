package persistence

import (
	"log"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ProvideDB(config *config.Config) *gorm.DB {
	return connect(config.Database.Url)
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		// users and permissions
		&models.Identity{},
		&models.Credential{},
		&models.RefreshToken{},
		&models.Invitation{},

		// application
		&models.Project{},
		&models.Board{},
		&models.Column{},
		&models.Task{},

		// monitoring and background tasks
		&models.Job{},
		&models.ActivityLog{},
	)

	if err != nil {
		return err
	}

	log.Printf("Database migrated")

	return nil
}

func connect(connectionString string) *gorm.DB {
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
