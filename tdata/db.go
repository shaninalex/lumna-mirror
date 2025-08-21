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
	database.ApplyMigrations(db)
	return db
}
