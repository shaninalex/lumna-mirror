// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"gitlab.com/shaninalex/flowreon/internal/web"
)

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		tokenService: web.NewTokenService(),
	}
}

type AuthHandler struct {
	tokenService *web.TokenService
}
