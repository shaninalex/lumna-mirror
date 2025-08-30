// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/org/adapter"
	"gitlab.com/shaninalex/flowreon/apps/org/domain"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// OrganizationHandler - organization handler.
type OrganizationHandler struct {
	manager domain.OrganizationManager
}

// NewOrganizationHandler - new organization handler.
func NewOrganizationHandler(manager domain.OrganizationManager) *OrganizationHandler {
	return &OrganizationHandler{
		manager: manager,
	}
}

// HandleGetByUser - handle get by user.
// TODO: rename
func (s *OrganizationHandler) HandleGetByUser(ctx *fiber.Ctx) error {
	organization, err := s.manager.Get(ctx.Context(), web.GetUserID(ctx))
	if errors.Is(err, apperrors.OrgNotFound) {
		return web.ReturnJSON(ctx, http.StatusNotFound, nil, err.Error())
	}
	return web.Success(ctx, adapter.ToDto(organization))
}
