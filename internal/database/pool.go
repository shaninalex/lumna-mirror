// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package database

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB - init db.
func InitDB(dsn string) *gorm.DB {
	var err error

	const maxRetries = 5
	const retryDelay = 5 * time.Second

	for i := 1; i <= maxRetries; i++ {
		gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		})
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
		return gormDB
	}

	log.Fatalf("[DB]: Failed to connect after %d attempts: %v", maxRetries, err)
	return nil
}
