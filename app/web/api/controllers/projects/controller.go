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

	router.GET("/projects", controller.List)
	router.POST("/projects", controller.Create)
	router.PATCH("/project/:id", controller.Patch)
	router.DELETE("/project/:id", controller.Delete)
}
