package project

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web"
)

type ProjectHandler struct {
	projectService services.ProjectManager
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{
		projectService: services.NewProjectManager(),
	}
}

func RegisterProjectController(router *web.Router) {
	h := NewProjectHandler()

	// list all user tokens
	router.GET("/api/v1/projects", h.List)
	router.POST("/api/v1/projects", h.Create)
	// s.router.GET("/api/v1/projects", projectHandler.List)
	// s.router.POST("/api/v1/projects", projectHandler.Create)
	// s.router.GET("/api/v1/project/{id}", projectHandler.Get)
	// s.router.DELETE("/api/v1/project/{id}", projectHandler.Delete)
	// s.router.PATCH("/api/v1/project/{id}", projectHandler.Patch)
}
