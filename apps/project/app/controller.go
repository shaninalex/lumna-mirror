package app

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
	"gitlab.com/shaninalex/jajirra/internal/pm"
	"gitlab.com/shaninalex/jajirra/internal/web"
)

func NewProjectController(router *fiber.App) {
	controller := ProjectController{
		router:     router,
		projectApi: pm.NewProjectManagement(),
	}
	controller.setRoutes()
}

type ProjectController struct {
	router     *fiber.App
	projectApi *pm.ProjectManagement
}

func (s *ProjectController) setRoutes() {
	s.router.Get("/api/project/list", s.HandleProjectsList)
	s.router.Get("/api/project/:projectKey/tasks", s.HandleProjectTasksList)
	s.router.Get("/api/project/:projectCode/tasks/:taskCode", s.HandleTaskDetail)
	s.router.Patch("/api/project/:projectCode/tasks/:taskCode", s.HandleTaskUpdate)
	s.router.Patch("/api/project/:projectKey/tasks/:taskID/status", s.HandleTaskPatchStatus)
}

func (s *ProjectController) HandleProjectsList(ctx *fiber.Ctx) error {
	// TODO: check user permission ( user should not see project he do not allowed to see )
	projects, err := s.projectApi.List(ctx.Context(), web.GetOrganizationId(ctx))
	if err != nil {
		if errors.Is(err, apperrors.ProjectNotFound) {
			return web.Success(ctx, NewProjectsDto(nil))
		}
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, NewProjectsDto(projects))
}

func (s *ProjectController) HandleProjectTasksList(ctx *fiber.Ctx) error {
	projectKey := ctx.Params("projectKey")
	statuses, err := s.projectApi.TasksList(ctx.Context(), web.GetOrganizationId(ctx), projectKey)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, NewTasksStatusDto(statuses))
}

func (s *ProjectController) HandleTaskPatchStatus(ctx *fiber.Ctx) error {
	in, err := NewPatchTaskInput(ctx)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	if err = s.projectApi.PatchTaskStatus(ctx.Context(), web.GetOrganizationId(ctx), in.ProjectKey, in.TaskID, in.Data); err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, nil, "Task saved")
}

func (s *ProjectController) HandleTaskDetail(ctx *fiber.Ctx) error {
	projectCode := ctx.Params("projectCode")
	taskCode := ctx.Params("taskCode")
	task, err := s.projectApi.TaskDetail(ctx.Context(), web.GetOrganizationId(ctx), projectCode, taskCode)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, NewTaskDto(task))
}

func (s *ProjectController) HandleTaskUpdate(ctx *fiber.Ctx) error {
	data, err := NewUpdateTaskInput(ctx)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	if err = s.projectApi.TaskUpdate(ctx.Context(), web.GetOrganizationId(ctx), data.ProjectCode, data.TaskCode, data.Data); err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, nil, "Task saved")
}

func (s *ProjectController) getFilterParams(ctx *fiber.Ctx) (*TaskFilter, error) {
	q := new(TaskFilter)
	err := ctx.QueryParser(q)
	return q, err
}
