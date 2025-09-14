// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon"
)

func HandleHealth(ctx *fiber.Ctx) error {
	ctx = ctx.Status(http.StatusOK)
	return ctx.JSON(map[string]interface{}{
		"status":  "ok",
		"name":    ctx.App().Config().AppName,
		"version": flowreon.Version,
	})
}
