// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/user/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
)

type UserHandler struct {
	manager domain.UserManager
}

// NewUserHandler - new user handler
func NewUserHandler(manager domain.UserManager) *UserHandler {
	return &UserHandler{
		manager: manager,
	}
}

// HandleGetUser - get user object
func (s *UserHandler) HandleGetUser(ctx *fiber.Ctx) error {
	user, err := s.manager.GetUser(ctx.Context(), web.GetUserID(ctx))
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, user)
}

// HandleUpdateSettings - update user settings
func (s *UserHandler) HandleUpdateSettings(ctx *fiber.Ctx) error {
	data, err := web.ParseBody[models.UserSettings](ctx)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	err = s.manager.UpdateUserSettings(ctx.Context(), web.GetUserID(ctx), data)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, nil, "Settings updated")
}
