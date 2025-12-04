package auth

import (
	"fmt"

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

func RegisterAuthController(router *web.Router, baseUrl string) {
	h := NewAuthHandler()
	router.POST(fmt.Sprintf("%s/login", baseUrl), h.HandleLogin)
	router.POST(fmt.Sprintf("%s/register", baseUrl), h.HandleRegistration)
	router.POST(fmt.Sprintf("%s/refresh", baseUrl), h.HandleRefresh)
}
