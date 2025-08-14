// Copyright © 2025 Soundstream https://soundstream.shaninalex.com. All rights reserved.

package db

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gorm.io/gorm"
)

type DBMiddleware struct {
	db *gorm.DB
}

func NewDBMiddleware(db *gorm.DB) fiber.Handler {
	m := &DBMiddleware{db: db}
	return m.Wrap()
}

func (m *DBMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals(base.ContextDB, m.db)
		return ctx.Next()
	}
}
