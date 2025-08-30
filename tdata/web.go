// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
	"gorm.io/gorm"
)

// AuthTestRouter - auth router.
func AuthTestRouter(db *gorm.DB, name string) *fiber.App {
	router := web.DefaultRouter(db, name)
	kratosClient := NewMockKratosService()
	router.Use(web.NewAuthMiddleware(kratosClient))
	return router
}

func SetAuthRequest(req *http.Request, user models.AuthUser, cookie string) {
	req.Header.Set("X-USER", user.GetID().String())
	req.Header.Set("Cookie", cookie)
}
