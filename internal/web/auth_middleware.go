// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
func NewAuthMiddleware(kratosService kratos.IKratos) fiber.Handler {
	m := &AuthMiddleware{
		kratosService: kratosService,
	}
	return m.Wrap()
}

// Wrap - wrap.
func (s *AuthMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Get("X-USER")
		if id == "" {
			return Error(ctx, http.StatusUnauthorized, fmt.Errorf("user is empty"))
		}

		userID, err := uuid.Parse(id)
		if err != nil {
			return Error(ctx, http.StatusUnauthorized, fmt.Errorf("user id is invalid"))
		}

		session, _, err := s.kratosService.GetSession(ctx.Context(), ctx.Get("cookie"))
		if err != nil {
			return Error(ctx, http.StatusUnauthorized, apperrors.SessionNotFound)
		}
		ctx.Locals(base.ContextSession, session)

		user := &models.User{ID: userID}
		db := database.GetDB(ctx.Context())
		tx := db.First(&user)
		if tx.Error != nil {
			return Error(ctx, http.StatusUnauthorized, apperrors.UserNotFound)
		}

		identity, _, err := s.kratosService.GetIdentity(ctx.Context(), id)
		if err != nil {
			return Error(ctx, http.StatusUnauthorized, apperrors.UserIdentityNotFound)
		}

		user.Identity = identity

		ctx.Locals(base.ContextUser, user)
		ctx.Locals(base.ContextUserID, userID)

		if user.OrganizationID == nil {
			return Error(ctx, http.StatusForbidden, apperrors.UserOrgNotAttached)
		}
		ctx.Locals(base.ContextOrgID, *user.OrganizationID)

		return ctx.Next()
	}
}
