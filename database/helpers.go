// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package database

import (
	"context"

	"gitlab.com/shaninalex/flowreon/internal/base"
	"gorm.io/gorm"
)

func GetDB(ctx context.Context) *gorm.DB {
	db := ctx.Value(base.ContextDB).(*gorm.DB)
	if db == nil {
		panic("postgres context is not set")
	}
	return db
}
