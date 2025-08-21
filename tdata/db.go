// Copyright © 2025 JaJirra https://jajirra.shaninalex.com. All rights reserved.

package tdata

import (
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

func clearDatabase() {
	tables := []any{
		&database.User{},
		&database.Issue{},
		&database.Epic{},
		&database.Sprint{},
		&database.Organization{},
		&database.Project{},
	}
	for _, t := range tables {
		// AllowGlobalUpdate ensures DELETE without WHERE is permitted
		if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(t).Error; err != nil {
			panic(err)
		}
	}
}
