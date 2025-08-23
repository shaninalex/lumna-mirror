package web

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/domain"
)

// GetKratosRedirectUrl return redirect with kratos base url from config
func GetKratosRedirectUrl(c base.IConfig, path string) string {
	return fmt.Sprintf("%s%s", c.String("kratos.url_browser"), path)
}

// ReturnJson return api response based on statuses
func ReturnJson(ctx *fiber.Ctx, status int, data any, params ...any) error {
	_data := domain.NewApiResponse(data)
	if status >= 400 {
		_data.Status = false
	}
	for _, p := range params {
		if msg, ok := p.(string); ok {
			_data.Messages = append(_data.Messages, msg)
		}
	}

	ctx.Status(status)
	return ctx.JSON(_data)
}

// Success return api response based on statuses
func Success(ctx *fiber.Ctx, data any, params ...any) error {
	return ReturnJson(ctx, http.StatusOK, data, params)
}

func GetUserId(ctx *fiber.Ctx) uuid.UUID {
	if id, ok := ctx.Locals(base.ContextUserID).(uuid.UUID); !ok {
		return id
	}
	return uuid.Nil
}

func GetOrganizationId(ctx *fiber.Ctx) uuid.UUID {
	if id, ok := ctx.Locals(base.ContextOrgID).(uuid.UUID); !ok {
		return id
	}
	return uuid.Nil
}

func GetUser(ctx *fiber.Ctx) *database.User {
	user, _ := ctx.Locals(base.ContextUser).(*database.User)
	return user
}

func GetSession(ctx *fiber.Ctx) *ory.Session {
	session, _ := ctx.Locals(base.ContextUserID).(*ory.Session)
	return session
}
