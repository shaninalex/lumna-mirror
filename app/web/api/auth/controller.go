package auth

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web"
)

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

type AuthHandler struct {
	userService services.UserManager
	authService services.AuthManager
}

func NewAuthController(router *web.Router) {
	h := NewAuthHandler()
	router.POST("/api/v1/auth/login", h.HandleLogin)
	router.POST("/api/v1/auth/register", h.HandleRegistration)
	router.POST("/api/v1/auth/refresh", h.HandleRefresh)
}
