// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// TaskHandler - task handler.
type TaskHandler struct {
	manager domain.ProjectManager
}

// NewTaskHandler - new task handler.
func NewTaskHandler(manager domain.ProjectManager) *TaskHandler {
	return &TaskHandler{
		manager: manager,
	}
}

// HandleTaskPatchStatus - handle task patch status.
func (s *TaskHandler) HandleTaskPatchStatus(ctx *fiber.Ctx) error {
	in, err := adapter.NewPatchTaskInput(ctx)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	if err = s.manager.PatchTaskStatus(ctx.Context(), web.GetOrganizationId(ctx), in.ProjectCode, in.TaskCode, in.Data); err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, nil, "Task saved")
}

// HandleTaskDetail - handle task detail.
func (s *TaskHandler) HandleTaskDetail(ctx *fiber.Ctx) error {
	projectCode := ctx.Params("projectCode")
	taskCode := ctx.Params("taskCode")
	task, err := s.manager.TaskDetail(ctx.Context(), web.GetOrganizationId(ctx), projectCode, taskCode)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, adapter.NewTaskDto(task))
}

// HandleTaskUpdate - handle task update.
func (s *TaskHandler) HandleTaskUpdate(ctx *fiber.Ctx) error {
	data, err := adapter.NewUpdateTaskInput(ctx)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	if err = s.manager.TaskUpdate(ctx.Context(), web.GetOrganizationId(ctx), data.ProjectCode, data.TaskCode, data.Data); err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, nil, "Task saved")
}
