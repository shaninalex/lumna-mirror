package app

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
	"gitlab.com/shaninalex/jajirra/internal/pm"
	"gitlab.com/shaninalex/jajirra/internal/web"
)

func NewTaskController(router *fiber.App) {
	controller := TaskController{
		router:     router,
		projectApi: pm.NewProjects(),
	}
	controller.setRoutes()
}

type TaskController struct {
	router     *fiber.App
	projectApi *pm.Projects
}

func (s *TaskController) setRoutes() {
	s.router.Get("/api/project/list", s.HandleProjectsList)
	s.router.Get("/api/project/tasks", s.HandleTasksList)
}

func (s *TaskController) HandleProjectsList(ctx *fiber.Ctx) error {
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

func (s *TaskController) HandleTasksList(ctx *fiber.Ctx) error {
	// TODO: check user permission. On error: 403 ( user should not get tasks of the project he do not allowed to get )
	q := new(TaskFilter)
	if err := ctx.QueryParser(q); err != nil {
		return web.Error(ctx, http.StatusBadRequest, errors.New("provide filter conditions"))
	}
	issues, err := s.projectApi.ProjectTasks(ctx.Context(), web.GetOrganizationId(ctx), q.Project)
	if err != nil {
		return web.Error(ctx, http.StatusBadRequest, err)
	}
	return web.Success(ctx, NewIssuesDto(issues))
}
