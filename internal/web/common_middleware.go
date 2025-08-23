package web

import (
	"github.com/gofiber/fiber/v2"
)

type CommonMiddleware struct {
}

func NewCommonMiddleware() fiber.Handler {
	m := &CommonMiddleware{}
	return m.Wrap()
}

func (m *CommonMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("ip", ctx.IP())
		// Other custom variables
		return ctx.Next()
	}
}
