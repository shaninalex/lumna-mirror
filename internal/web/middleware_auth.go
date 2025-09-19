// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gitlab.com/shaninalex/flowreon/models"
)

// AuthMiddleware - auth middleware.
type AuthMiddleware struct {
	kratosService kratos.IKratos
}

// NewAuthMiddleware - new auth middleware.
func NewAuthMiddleware(kratosService kratos.IKratos) *AuthMiddleware {
	return &AuthMiddleware{
		kratosService: kratosService,
	}
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

		session, _, err := s.kratosService.GetSession(ctx, r.Header.Get("cookie"))
		if err != nil {
			Error(w, http.StatusUnauthorized, apperrors.SessionNotFound)
			return
		}
		ctx = context.WithValue(ctx, base.ContextSession, session)

		user := &models.User{ID: userID}
		tx := database.GetDB(ctx).Preload("Organization").First(&user)
		if tx.Error != nil {
			Error(w, http.StatusUnauthorized, apperrors.UserNotFound)
			return
		}

		identity, _, err := s.kratosService.GetIdentity(ctx, id)
		if err != nil {
			Error(w, http.StatusUnauthorized, apperrors.UserIdentityNotFound)
			return
		}

		user.Identity = identity

		ctx = context.WithValue(ctx, base.ContextUser, user)
		ctx = context.WithValue(ctx, base.ContextUserID, userID)

		if user.OrganizationID == nil {
			Error(w, http.StatusForbidden, apperrors.UserOrgNotAttached)
			return
		}
		ctx = context.WithValue(ctx, base.ContextOrgID, *user.OrganizationID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
