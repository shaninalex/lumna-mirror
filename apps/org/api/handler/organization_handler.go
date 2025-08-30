// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/org/adapter"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

func (s *OrganizationHandler) HandleGetByUser(ctx *fiber.Ctx) error {
	organization, err := s.api.Get(ctx.Context(), web.GetUserId(ctx))
	if errors.Is(err, apperrors.OrgNotFound) {
		return web.ReturnJson(ctx, http.StatusNotFound, nil, err.Error())
	}
	return web.Success(ctx, adapter.ToDto(organization))
}

//func (s *OrganizationHandler) HandleCreate(ctx *fiber.Ctx) error {
//	return nil
//}
//
//func (s *OrganizationHandler) HandlePatch(ctx *fiber.Ctx) error {
//	return nil
//}
