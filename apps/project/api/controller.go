// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"gitlab.com/shaninalex/flowreon/apps/project/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
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
	projectHandler := handler.NewProjectHandler(domain.NewProjectManagement())
	s.router.GET("/api/project", projectHandler.HandleProjectsList)
	s.router.POST("/api/project", projectHandler.HandleProjectCreate)

	taskHandler := handler.NewTaskHandler(domain.NewProjectManagement())
	s.router.POST("/api/project/tasks", taskHandler.HandleTaskCreate)
	s.router.GET("/api/project/{projectCode}/tasks", taskHandler.HandleProjectTasksList)
	s.router.GET("/api/project/{projectCode}/tasks/{taskCode}", taskHandler.HandleTaskDetail)
	s.router.PATCH("/api/project/{projectCode}/tasks/{taskCode}", taskHandler.HandleTaskUpdate)
	s.router.PATCH("/api/project/{projectCode}/tasks/{taskCode}/status", taskHandler.HandleTaskPatchStatus)
}
