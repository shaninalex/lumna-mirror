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

	// Retrieve all projects
	s.router.GET("/api/v1/projects", projectHandler.HandleProjectsList)
	// Create a new project
	s.router.POST("/api/v1/projects", projectHandler.HandleProjectCreate)

	taskHandler := handler.NewTaskHandler(domain.NewProjectManagement())
	s.router.POST("/api/v1/projects/tasks", taskHandler.HandleTaskCreate)
	s.router.GET("/api/v1/projects/{projectCode}/tasks", taskHandler.HandleProjectTasksList)
	s.router.GET("/api/v1/projects/{projectCode}/tasks/{taskCode}", taskHandler.HandleTaskDetail)
	s.router.PATCH("/api/v1/projects/{projectCode}/tasks/{taskCode}", taskHandler.HandleTaskUpdate)
	s.router.PATCH("/api/v1/projects/{projectCode}/tasks/{taskCode}/status", taskHandler.HandleTaskPatchStatus)

	/*

		GET
		/api/v1/projects/{id}
		Retrieve a specific project

		DELETE
		/api/v1/projects/{id}
		Delete Project

		PATCH
		/api/v1/projects/{id}
		Update specific project

		GET
		/api/v1/projects/{id}/tasks
		Retrieve tasks for a project

		POST
		/api/v1/projects/{id}/tasks
		Create a new task in a project

		GET
		/api/v1/project/{id}/badges
		Retrieve all badges

		POST
		/api/v1/project/{id}/badges
		Create project badge

		DELETE
		/api/v1/project/{id}/badges/{badgeId}
		Delete project badge










	*/

}
