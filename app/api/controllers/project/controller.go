package project

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type ProjectController struct {
	projectService *services.ProjectService
	boardService   *services.BoardService
}

func NewProjectsController(
	projectService *services.ProjectService,
	boardService *services.BoardService,
) *ProjectController {
	s := &ProjectController{
		projectService: projectService,
		boardService:   boardService,
	}
	return s
}

func (s *ProjectController) Register(router *gin.RouterGroup) {
	router.GET("/projects", s.List)
	router.POST("/projects", s.Create)
	router.PATCH("/project/:id", s.Patch)
	router.DELETE("/project/:id", s.Delete)
}
