package app

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
	s.router.Get("/api/project/:projectId/tasks", s.HandleProjectTasks)
}

func (s *TaskController) HandleProjectsList(ctx *fiber.Ctx) error {
	id := web.GetUserId(ctx)
	projects, err := s.projectApi.List(ctx.Context(), id)
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}
	return web.Success(ctx, NewProjectsDto(projects))
}

func (s *TaskController) HandleProjectTasks(ctx *fiber.Ctx) error {
	return web.ReturnJson(ctx, http.StatusNotImplemented, nil, fmt.Errorf("endpoint is not implemented yet"))
}
