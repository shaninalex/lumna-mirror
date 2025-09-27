// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package auth

import (
	"gitlab.com/shaninalex/flowreon/apps/auth/handlers"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

func NewAuthController(router *web.Router) {
	h := handlers.NewAuthHandler()
	router.POST("/api/v1/auth/login", h.HandleLogin)
	router.POST("/api/v1/auth/register", h.HandleRegistration)
	router.POST("/api/v1/auth/refresh", h.HandleRefresh)
}
