// Copyright © 2025 Lumna. All rights reserved.

package api

import (
	"gitlab.com/shaninalex/flowreon/apps/project/api/handler"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// ProjectController - project controller.
type ProjectController struct {
	router *web.Router
}

// NewProjectController - new project controller.
func NewProjectController(router *web.Router) {
	controller := ProjectController{router: router}
	controller.init()
}

func (s *ProjectController) init() {
	// ProjectHandler
	projectHandler := handler.NewProjectHandler()
	s.router.GET("/api/v1/projects", projectHandler.List)
	s.router.POST("/api/v1/projects", projectHandler.Create)
	s.router.GET("/api/v1/project/{id}", projectHandler.Get)
	s.router.DELETE("/api/v1/project/{id}", projectHandler.Delete)
	s.router.PATCH("/api/v1/project/{id}", projectHandler.Patch)

	// ProjectStatusHandler
	projectStatusHandler := handler.NewProjectStatusHandler()
	s.router.GET("/api/v1/project/{id}/statuses", projectStatusHandler.Get)
	s.router.POST("/api/v1/project/{id}/statuses", projectStatusHandler.Post)
	s.router.PATCH("/api/v1/project/{id}/statuses-sort", projectStatusHandler.PatchSort)
	s.router.PATCH("/api/v1/project/{id}/statuses/{statusId}", projectStatusHandler.Patch)
	s.router.DELETE("/api/v1/project/{id}/statuses/{statusId}", projectStatusHandler.Delete)

	// ProjectTaskHandler
	projectTaskHandler := handler.NewProjectTaskHandler()
	s.router.GET("/api/v1/project/{id}/tasks", projectTaskHandler.List)
	s.router.POST("/api/v1/project/{id}/tasks", projectTaskHandler.Create)

	// ProjectBadgeHandler
	projectBadgeHandler := handler.NewProjectBadgeHandler()
	s.router.GET("/api/v1/project/{id}/badges", projectBadgeHandler.List)
	s.router.POST("/api/v1/project/{id}/badges", projectBadgeHandler.Create)
	s.router.DELETE("/api/v1/project/{id}/badges/{badgeId}", projectBadgeHandler.Delete)
}
