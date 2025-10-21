// Copyright © 2025 Lumna. All rights reserved.

package handlers

import (
	"github.com/shaninalex/lumna/app/internal/token"
)

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		authService: token.NewAuthService(),
	}
}

type AuthHandler struct {
	authService token.ApiAuthService
}
