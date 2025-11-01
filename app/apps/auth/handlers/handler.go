package handlers

import (
	"gitlab.com/shaninalex/lumna/app/domain"
	"gitlab.com/shaninalex/lumna/app/internal/token"
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
