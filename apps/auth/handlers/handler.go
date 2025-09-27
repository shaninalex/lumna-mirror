// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handlers

import (
	"gitlab.com/shaninalex/flowreon/internal/token"
)

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		tokenManager: token.NewTokenManager(),
	}
}

type AuthHandler struct {
	tokenManager token.TokenManager
}
