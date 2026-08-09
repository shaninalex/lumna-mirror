package project

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type ProjectController struct {
	projectService services.ProjectService
	boardService   services.ListService
}

func NewProjectsController(
	projectService services.ProjectService,
	boardService services.ListService,
) *ProjectController {
	s := &ProjectController{
		projectService: projectService,
		boardService:   boardService,
	}
	return s
}

func (s *ProjectController) Register(router *gin.RouterGroup) {
	router.GET("/projects/:workspaceID", s.List)
	router.GET("/projects/generate-key", s.GenerateKey)
	router.POST("/projects", s.Create)
	router.PATCH("/projects/:id", s.Patch)
	router.DELETE("/projects/:id", s.Delete)
}
