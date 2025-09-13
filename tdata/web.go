// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
)

// AuthTestRouter - auth router.
func AuthTestRouter(ctx context.Context) *fiber.App {
	db := database.GetDB(ctx)

	router := web.DefaultRouter(db, "test_router")
	kratosClient := NewMockKratosService()
	router.Use(web.NewAuthMiddleware(kratosClient))
	return router
}

// SetAuthRequest - sets the auth request.
func SetAuthRequest(req *http.Request, user models.AuthUser, cookie string) {
	req.Header.Set("X-USER", user.GetID().String())
	req.Header.Set("Cookie", cookie)
}
