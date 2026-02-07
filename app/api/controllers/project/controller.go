package project

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/internal/logger"
	"gitlab.com/shaninalex/lumna/app/services"
)

type ProjectController struct {
	projectService *services.ProjectService
	boardService   *services.BoardService
	logger         logger.Logger
}

func NewProjectsController(
	projectService *services.ProjectService,
	boardService *services.BoardService,
	logger logger.Logger,
) *ProjectController {
	s := &ProjectController{
		projectService: projectService,
		boardService:   boardService,
		logger:         logger,
	}

	return s
}

func (s *ProjectController) Register(router *gin.RouterGroup) {
	router.GET("/projects", s.List)
	router.POST("/projects", s.Create)
	router.PATCH("/project/:id", s.Patch)
	router.DELETE("/project/:id", s.Delete)
}
