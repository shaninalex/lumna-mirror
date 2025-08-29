// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"

	"gitlab.com/shaninalex/flowreon/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitTestDatabase() *gorm.DB {
	var err error
	db, err = gorm.Open(postgres.Open(config.String("app.dsn")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}

	db.Exec("CREATE SCHEMA IF NOT EXISTS kratos")
	db.Exec("CREATE SCHEMA IF NOT EXISTS keto")
	db.Exec("CREATE SCHEMA IF NOT EXISTS testing")

	db.Exec("SET search_path='testing'")

	database.ApplyMigrations(db)
	return db
}

func clearDatabase(ctx context.Context) {
	_db := database.GetDB(ctx)
	if err := _db.Exec(`
		TRUNCATE TABLE 
			users, issues, epics, sprints, organizations, projects 
		RESTART IDENTITY CASCADE
	`).Error; err != nil {
		panic(err)
	}
}
