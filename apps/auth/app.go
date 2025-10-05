// Copyright © 2025 Lumna. All rights reserved.

package auth

import (
	"github.com/shaninalex/lumna/apps/auth/handlers"
	"github.com/shaninalex/lumna/internal/web"
)

func NewAuthController(router *web.Router) {
	h := handlers.NewAuthHandler()
	router.POST("/api/v1/auth/login", h.HandleLogin)
	router.POST("/api/v1/auth/register", h.HandleRegistration)
	router.POST("/api/v1/auth/refresh", h.HandleRefresh)
}
