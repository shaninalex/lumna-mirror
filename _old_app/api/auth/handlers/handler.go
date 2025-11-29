package handlers

import (
	"gitlab.com/shaninalex/lumna/_old_app/domain"
	"gitlab.com/shaninalex/lumna/_old_app/pkg/token"
)

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: token.NewAuthService(),
		userService: domain.NewUserService(),
	}
}

type AuthHandler struct {
	authService token.ApiAuthService
	userService *domain.UserService
}
