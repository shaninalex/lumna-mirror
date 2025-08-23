// Copyright © 2025 Jajirra https://jajirra.shaninalex.com. All rights reserved.

package database

import (
	"log"
	"time"

	"gitlab.com/shaninalex/jajirra/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
	var err error

	const maxRetries = 5
	const retryDelay = 5 * time.Second

	for i := 1; i <= maxRetries; i++ {
		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("[DB]: Connection attempt %d/%d failed to open: %v", i, maxRetries, err)
			time.Sleep(retryDelay)
			continue
		}

		sqlDB, err := gormDB.DB()
		if err != nil {
			log.Printf("[DB]: Connection attempt %d/%d sql db is not valid: %v", i, maxRetries, err)
			time.Sleep(retryDelay)
			continue
		}

		err = sqlDB.Ping()
		if err != nil {
			log.Printf("[DB]: Connection attempt %d/%d failed ping sql db: %v", i, maxRetries, err)
			time.Sleep(retryDelay)
			continue
		}

		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour)
		ApplyMigrations(gormDB)
		return gormDB
	}

	log.Fatalf("[DB]: Failed to connect after %d attempts: %v", maxRetries, err)
	return nil
}

// ApplyMigrations applying migrations
func ApplyMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.User{},
		&models.Issue{},
		&models.IssueStatus{},
		&models.Epic{},
		&models.Sprint{},
		&models.Organization{},
		&models.Project{},
	)
	if err != nil {
		log.Printf("[DB]: Unable to apply migrations: %v", err)
		panic(err)
	}
}
