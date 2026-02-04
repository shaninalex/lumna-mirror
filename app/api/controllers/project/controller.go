package project

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type ProjectController struct {
	projectService *services.ProjectService
	boardService   *services.BoardService
}

func NewProjectsController() *ProjectController {
	s := &ProjectController{
		projectService: services.NewProjectService(),
		boardService:   services.NewBoardService(),
	}

	return s
}

func RegisterProjectController(router *gin.RouterGroup) {
	controller := NewProjectsController()

	router.GET("/projects", controller.List)
	router.POST("/projects", controller.Create)
	router.PATCH("/project/:id", controller.Patch)
	router.DELETE("/project/:id", controller.Delete)
}
