package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

type CommmonMiddleware struct {
}

func NewCommmonMiddleware() fiber.Handler {
	m := &CommmonMiddleware{}
	return m.Wrap()
}

func (m *CommmonMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("ip", ctx.IP())
		// Other custom variables
		return ctx.Next()
	}
}
