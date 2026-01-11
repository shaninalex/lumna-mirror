package auth

import (
	"gitlab.com/shaninalex/lumna/app2/services"
	"gitlab.com/shaninalex/lumna/app2/web"
)

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userService: services.NewUserManager(),
		authService: services.NewAuthManager(),
	}
}

type AuthHandler struct {
	userService services.UserManager
	authService services.AuthManager
}

func RegisterAuthController(router *web.Router) {
	h := NewAuthHandler()
	router.POST("/api/v1/auth/login", h.HandleLogin)
	router.POST("/api/v1/auth/register", h.HandleRegistration)
	router.GET("/api/v1/auth/refresh", h.HandleRefresh)
}
