package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/web"
	"gorm.io/gorm"
)

func NewTaskController(
	config base.IConfig,
	router *fiber.App,
) {
	controller := TaskController{
		config: config,
		router: router,
	}
	controller.setRoutes()
}

type TaskController struct {
	router            *fiber.App
	config            base.IConfig
	projectRepository *database.ProjectRepository
}

func (s *TaskController) setRoutes() {
	s.router.Get("/api/project/list", s.HandleProjectsList)
}

func (s *TaskController) HandleProjectsList(ctx *fiber.Ctx) error {
	id := web.GetUserId(ctx)
	projects, err := s.projectRepository.List(ctx.Context(), func(q *gorm.DB) *gorm.DB {
		return q.Where("user_id = ?", id)
	})
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}
	return web.Success(ctx, projects)
}
