package web

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/kratos"
	"gorm.io/gorm"
)

type AuthMiddleware struct {
	kratosService  kratos.IKratos
	userRepository *database.UserRepository
}

func NewAuthMiddleware(kratosService kratos.IKratos) fiber.Handler {
	m := &AuthMiddleware{
		kratosService:  kratosService,
		userRepository: database.NewUserRepository(),
	}
	return m.Wrap()
}

func (s *AuthMiddleware) Wrap() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Get("X-USER")
		if id == "" {
			return ReturnJson(ctx, http.StatusUnauthorized, nil, "User is empty")
		}
		userID, err := uuid.Parse(id)
		if err != nil {
			return ReturnJson(ctx, http.StatusUnauthorized, nil, "User id is invalid")
		}

		session, _, err := s.kratosService.GetSession(ctx.Context(), ctx.Get("cookie"))
		if err != nil {
			return ReturnJson(ctx, http.StatusUnauthorized, nil, "Session not set")
		}
		ctx.Locals(base.ContextSession, session)

		user, err := s.userRepository.GetByID(ctx.Context(), userID)
		if err != nil {
			return ReturnJson(ctx, http.StatusUnauthorized, nil, "User not found")
		}

		identity, _, err := s.kratosService.GetIdentity(ctx.Context(), id)
		if err != nil {
			return ReturnJson(ctx, http.StatusUnauthorized, nil, "Identity not found")
		}

		user.Identity = identity

		ctx.Locals(base.ContextUser, user)
		ctx.Locals(base.ContextUserID, userID)
		db := database.GetDB(ctx.Context())

		orgID, err := gorm.G[string](db).Raw("SELECT id FROM organizations WHERE user_id = ?", userID).Find(ctx.Context())
		if err != nil {
			log.Error("user does not attach to any organizations")
		}
		ctx.Locals(base.ContextOrgID, orgID)

		return ctx.Next()
	}
}
