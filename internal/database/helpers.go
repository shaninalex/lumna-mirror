// Copyright © 2025 Lumna. All rights reserved.

package database

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/flowreon/internal/base"
)

// GetDb - returns the *sql.DB
func GetDb(ctx context.Context) *sql.DB {
	sqlDB := ctx.Value(base.ContextDB).(*sql.DB)
	if sqlDB == nil {
		panic("postgres context is not set")
	}
	return sqlDB
}
