// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models/repositories"
)

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		tokenService:   web.NewTokenService(),
		userRepository: repositories.NewUserRepository(),
	}
}

type AuthHandler struct {
	tokenService   *web.TokenService
	userRepository *repositories.UserRepository
}
