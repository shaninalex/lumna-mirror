// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package database

import (
	"context"
	"database/sql"
	"fmt"

	"gitlab.com/shaninalex/flowreon/internal/base"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDB - returns the db.
func GetDB(ctx context.Context) *gorm.DB {
	sqlDB := ctx.Value(base.ContextDB).(*sql.DB)
	if sqlDB == nil {
		panic("postgres context is not set")
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create gorm.DB: %v", err))
	}
	return db
}

// GetDb - returns the *sql.DB
func GetDb(ctx context.Context) *sql.DB {
	sqlDB := ctx.Value(base.ContextDB).(*sql.DB)
	if sqlDB == nil {
		panic("postgres context is not set")
	}
	return sqlDB
}
