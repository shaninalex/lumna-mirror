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
	s.router.Get("/api/project/:projectKey/statuses", s.HandleProjectStatuses)
	s.router.Get("/api/project/tasks", s.HandleTasksList)
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

func (s *ProjectController) HandleTasksList(ctx *fiber.Ctx) error {
	// TODO: check user permission. On error: 403 ( user should not get tasks of the project he do not allowed to get )
	q, err := s.getFilterParams(ctx)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, errors.New("provide proper filter conditions"))
	}
	issues, err := s.projectApi.Issues(ctx.Context(), web.GetOrganizationId(ctx), q.Project)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, NewTasksDto(issues))
}

func (s *ProjectController) HandleProjectStatuses(ctx *fiber.Ctx) error {
	projectKey := ctx.Params("projectKey")
	statuses, err := s.projectApi.Statuses(ctx.Context(), web.GetOrganizationId(ctx), projectKey)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, NewTasksStatusDto(statuses))
}

func (s *ProjectController) getFilterParams(ctx *fiber.Ctx) (*TaskFilter, error) {
	q := new(TaskFilter)
	err := ctx.QueryParser(q)
	return q, err
}
