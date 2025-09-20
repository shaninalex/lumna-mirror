// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package auth

import (
	"gitlab.com/shaninalex/flowreon/apps/auth/handlers"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

func NewAuthController(router *web.Router) {
	router.GET("/auth/login", handlers.HandleLoginTemplate)
	router.POST("/auth/login", handlers.HandleLogin)
}
