// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/auth/domain"
	"gitlab.com/shaninalex/flowreon/apps/auth/dto"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

type AuthHooksHandler struct {
	authApi domain.AuthHookHandler
}

func NewAuthHooksHandler(api domain.AuthHookHandler) *AuthHooksHandler {
	// hooks
	// TODO: we can authenticate hooks by "ory_kratos_continuity" cookie
	return &AuthHooksHandler{
		authApi: api,
	}
}

func (s *AuthHooksHandler) HandleHookRegister(ctx *fiber.Ctx) error {
	var data dto.HooksKratosPayloadDTO
	err := ctx.BodyParser(&data)
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}
	err = s.authApi.HookRegister(ctx.Context(), &data)
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}
	return web.Success(ctx, nil)
}

func (s *AuthHooksHandler) HandleHookVerify(ctx *fiber.Ctx) error {
	return web.Success(ctx, nil)
}

func (s *AuthHooksHandler) HandleHookLogin(ctx *fiber.Ctx) error {
	return web.Success(ctx, nil)
}
