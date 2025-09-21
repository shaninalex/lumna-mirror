// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/models"
)

// AuthMiddleware - auth middleware.
type AuthMiddleware struct {
}

// NewAuthMiddleware - new auth middleware.
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

// Wrap - wrap.
func (s *AuthMiddleware) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.Header.Get("X-USER")
		if id == "" {
			Error(w, http.StatusUnauthorized, fmt.Errorf("user is empty"))
			return
		}

		userID, err := uuid.Parse(id)
		if err != nil {
			Error(w, http.StatusUnauthorized, fmt.Errorf("user id is invalid"))
			return
		}

		user := &models.User{ID: userID}
		tx := database.GetDB(ctx).Preload("Organization").First(&user)
		if tx.Error != nil {
			Error(w, http.StatusUnauthorized, apperrors.UserNotFound)
			return
		}

		ctx = context.WithValue(ctx, base.ContextUser, user)
		ctx = context.WithValue(ctx, base.ContextUserID, userID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
