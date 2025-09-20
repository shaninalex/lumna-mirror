// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models/repositories"
)

func NewAuthHandler(sessionStore *web.CookieStoreDatabase) *AuthHandler {
	return &AuthHandler{
		sessionStore:   sessionStore,
		userRepository: repositories.NewUserRepository(),
	}
}

type AuthHandler struct {
	sessionStore   *web.CookieStoreDatabase
	userRepository *repositories.UserRepository
}
