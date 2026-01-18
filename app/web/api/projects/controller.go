package projects

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/shaninalex/lumna/app/services"
)

type ProjectsController struct {
	projectService *services.ProjectService
}

func NewProjectsController() *ProjectsController {
	s := &ProjectsController{
		projectService: services.NewProjectService(),
	}

	return s
}

func NewController(router *gin.RouterGroup) {
	controller := NewProjectsController()

	router.GET("/projects", controller.GetProjects)
}
