// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

type ProjectHandler struct {
	manager domain.ProjectManager
}

func NewProjectHandler(manager domain.ProjectManager) *ProjectHandler {
	h := &ProjectHandler{
		manager: manager,
	}
	return h
}

func (s *ProjectHandler) HandleProjectsList(ctx *fiber.Ctx) error {
	// TODO: check user permission ( user should not see project he do not allowed to see )
	projects, err := s.manager.List(ctx.Context(), web.GetOrganizationId(ctx))
	if err != nil {
		if errors.Is(err, apperrors.ProjectNotFound) {
			return web.Success(ctx, adapter.NewProjectsDto(nil))
		}
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, adapter.NewProjectsDto(projects))
}

func (s *ProjectHandler) HandleProjectTasksList(ctx *fiber.Ctx) error {
	projectCode := ctx.Params("projectCode")
	statuses, err := s.manager.TasksList(ctx.Context(), web.GetOrganizationId(ctx), projectCode)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, adapter.NewTasksStatusDto(statuses))
}

type TaskFilter struct {
	Project  string `query:"project,required"`
	TaskCode string `query:"taskCode"`
}

func (s *ProjectHandler) getFilterParams(ctx *fiber.Ctx) (*TaskFilter, error) {
	q := new(TaskFilter)
	err := ctx.QueryParser(q)
	return q, err
}
