// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package web

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/models"
)

// GetKratosRedirectURL return redirect with kratos base url from config
func GetKratosRedirectURL(c base.IConfig, path string) string {
	return fmt.Sprintf("%s%s", c.String("kratos.url_browser"), path)
}

// ReturnJSON return api response based on statuses
func ReturnJSON(ctx *fiber.Ctx, status int, data any, params ...any) error {
	resp := NewAPIResponse(data)
	if status >= 400 {
		resp.Status = false
	}

	// separate messages vs app errors
	for _, p := range params {
		switch v := p.(type) {
		case string:
			// general user-facing message
			resp.Messages = append(resp.Messages, v)
		case apperrors.AppError:
			// structured app error
			resp.Errors = append(resp.Errors, v)
		case error:
			// try to extract AppError if wrapped
			var appErr apperrors.AppError
			if errors.As(v, &appErr) {
				resp.Errors = append(resp.Errors, appErr)
			} else {
				// fallback: put error string in messages
				resp.Messages = append(resp.Messages, v.Error())
			}
		}
	}

	ctx.Status(status)
	return ctx.JSON(resp)
}

// Success return api response based on statuses
func Success(ctx *fiber.Ctx, data any, params ...any) error {
	return ReturnJSON(ctx, http.StatusOK, data, params...)
}

// Error return api response based on statuses
func Error(ctx *fiber.Ctx, status int, err error) error {
	return ReturnJSON(ctx, status, nil, err)
}

// GetUserID - returns the user id.
func GetUserID(ctx *fiber.Ctx) uuid.UUID {
	if id, ok := ctx.Locals(base.ContextUserID).(uuid.UUID); ok {
		return id
	}
	return uuid.Nil
}

// GetOrganizationID - returns the organization id.
func GetOrganizationID(ctx *fiber.Ctx) uuid.UUID {
	if id, ok := ctx.Locals(base.ContextOrgID).(uuid.UUID); ok {
		return id
	}
	return uuid.Nil
}

// GetUser - returns the user.
func GetUser(ctx *fiber.Ctx) *models.User {
	user, _ := ctx.Locals(base.ContextUser).(*models.User)
	return user
}

// GetSession - returns the session.
func GetSession(ctx *fiber.Ctx) *ory.Session {
	session, _ := ctx.Locals(base.ContextUserID).(*ory.Session)
	return session
}

// ParamUUID - param uuid.
func ParamUUID(ctx *fiber.Ctx, name string) (uuid.UUID, error) {
	val := ctx.Params(name)
	id, err := uuid.Parse(val)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s is not a valid UUID", name)
	}
	return id, nil
}

// ParamString - param string.
func ParamString(ctx *fiber.Ctx, name string) (string, error) {
	val := ctx.Params(name)
	if val == "" {
		return "", fmt.Errorf("%s is required", name)
	}
	return val, nil
}

// ParseBody - parse fiber post/patch/put request to defined structure
func ParseBody[T any](ctx *fiber.Ctx) (*T, error) {
	var dto T
	if err := ctx.BodyParser(&dto); err != nil {
		return nil, err
	}
	return &dto, nil
}
