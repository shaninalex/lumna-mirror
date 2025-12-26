package project

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/web"
)

type ProjectHandler struct {
	projectService services.ProjectManager
	boardService   *services.BoardService
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{
		projectService: services.NewProjectManager(),
		boardService:   services.NewBoardService(),
	}
}

func RegisterProjectController(router *web.Router) {
	h := NewProjectHandler()

	// list all user tokens
	router.GET("/api/v1/projects", h.List)
	router.POST("/api/v1/projects", h.Create)
	// s.router.GET("/api/v1/project/{id}", h.Get)
	router.DELETE("/api/v1/project/{id}", h.Delete)
	// s.router.PATCH("/api/v1/project/{id}", h.Patch)

	router.GET("/api/v1/project/{id}/boards", h.BoardsList)
	router.POST("/api/v1/project/{id}/boards", h.BoardsCreate)
}
