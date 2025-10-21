// Copyright © 2025 Lumna. All rights reserved.

package db

import (
	"context"
	"database/sql"

	"github.com/shaninalex/lumna/app/internal/base"
)

// GetDb - returns the *sql.DB
func GetDb(ctx context.Context) *sql.DB {
	sqlDB := ctx.Value(base.ContextDB).(*sql.DB)
	if sqlDB == nil {
		panic("postgres context is not set")
	}
	return sqlDB
}
