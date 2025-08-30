// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"github.com/gofiber/fiber/v2"
)

// CommonMiddleware - common middleware.
type CommonMiddleware struct {
}

// NewCommonMiddleware - new common middleware.
func NewCommonMiddleware() fiber.Handler {
	m := &CommonMiddleware{}
	return m.Wrap()
}

// Wrap - wrap.
func (m *CommonMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("ip", ctx.IP())
		// Other custom variables
		return ctx.Next()
	}
}
