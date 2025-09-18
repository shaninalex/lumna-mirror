// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package database

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gorm.io/gorm"
)

// DbMiddleware - db middleware.
type DbMiddleware struct {
	db *gorm.DB
}

// NewDbMiddleware - new db middleware.
func NewDbMiddleware(db *gorm.DB) fiber.Handler {
	m := &DbMiddleware{db: db}
	return m.Wrap()
}

// Wrap - wrap.
func (m *DbMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals(base.ContextDB, m.db)
		return ctx.Next()
	}
}

type Middleware struct {
	db *sql.DB
}

func NewMiddleware(db *sql.DB) *Middleware {
	return &Middleware{db: db}
}

func (s *Middleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r = r.WithContext(context.WithValue(r.Context(), base.ContextDB, s.db))
		next.ServeHTTP(w, r)
	})
}
