package db

import (
	"context"

	"gitlab.com/shaninalex/jajirra/internal/base"
	"gorm.io/gorm"
)

func GetDB(ctx context.Context) *gorm.DB {
	db := ctx.Value(base.ContextDB).(*gorm.DB)
	if db == nil {
		panic("postgres context is not set")
	}
	return db
}
