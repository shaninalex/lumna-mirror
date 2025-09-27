// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"gitlab.com/shaninalex/flowreon/internal/token"
)

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: token.NewAuthService(),
	}
}

type AuthHandler struct {
	authService token.ApiAuthService
}
