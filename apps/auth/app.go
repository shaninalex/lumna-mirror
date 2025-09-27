// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package auth

import (
	"gitlab.com/shaninalex/flowreon/apps/auth/handlers"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

func NewAuthController(router *web.Router) {
	h := handlers.NewAuthHandler()
	router.POST("/api/auth/login", h.HandleLogin)
	router.POST("/api/auth/register", h.HandleRegistration)
	router.POST("/api/auth/refresh-token", h.HandleRefresh)
}
