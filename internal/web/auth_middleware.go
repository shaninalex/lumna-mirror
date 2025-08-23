package web

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/kratos"
	"gitlab.com/shaninalex/jajirra/models"
)

type AuthMiddleware struct {
	kratosService kratos.IKratos
}

func NewAuthMiddleware(kratosService kratos.IKratos) fiber.Handler {
	m := &AuthMiddleware{
		kratosService: kratosService,
	}
	return m.Wrap()
}

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
