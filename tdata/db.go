// Copyright © 2025 JaJirra https://jajirra.shaninalex.com. All rights reserved.

package tdata

import (
	"context"

	"gitlab.com/shaninalex/jajirra/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitTestDatabase() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	db.Exec("CREATE SCHEMA IF NOT EXISTS kratos")
	db.Exec("CREATE SCHEMA IF NOT EXISTS keto")
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
