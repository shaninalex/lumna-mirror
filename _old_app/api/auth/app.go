package auth

import (
	"gitlab.com/shaninalex/lumna/_old_app/apps/auth/handlers"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/web"
)

func NewAuthController(router *web.Router) {
	h := handlers.NewAuthHandler()
	router.POST("/api/v1/auth/login", h.HandleLogin)
	router.POST("/api/v1/auth/register", h.HandleRegistration)
	router.POST("/api/v1/auth/refresh", h.HandleRefresh)
}
