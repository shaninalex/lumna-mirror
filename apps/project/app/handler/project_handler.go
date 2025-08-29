// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

func (s *ProjectHandler) HandleProjectsList(ctx *fiber.Ctx) error {
	// TODO: check user permission ( user should not see project he do not allowed to see )
	projects, err := s.projectApi.List(ctx.Context(), web.GetOrganizationId(ctx))
	if err != nil {
		if errors.Is(err, apperrors.ProjectNotFound) {
			return web.Success(ctx, dto.NewProjectsDto(nil))
		}
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, dto.NewProjectsDto(projects))
}

func (s *ProjectHandler) HandleProjectTasksList(ctx *fiber.Ctx) error {
	projectCode := ctx.Params("projectCode")
	statuses, err := s.projectApi.TasksList(ctx.Context(), web.GetOrganizationId(ctx), projectCode)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, dto.NewTasksStatusDto(statuses))
}
