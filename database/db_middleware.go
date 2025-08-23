// Copyright © 2025 Jajirra https://jajirra.shaninalex.com. All rights reserved.

package database

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gorm.io/gorm"
)

type DbMiddleware struct {
	db *gorm.DB
}

func NewDbMiddleware(db *gorm.DB) fiber.Handler {
	m := &DbMiddleware{db: db}
	return m.Wrap()
}

func (m *DbMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals(base.ContextDB, m.db)
		return ctx.Next()
	}
}
