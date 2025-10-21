// Copyright © 2025 Lumna. All rights reserved.

package api

import (
	"github.com/shaninalex/lumna/app/apps/task/api/handlers"
	"github.com/shaninalex/lumna/app/internal/web"
)

// TaskController - task controller.
type TaskController struct {
	router *web.Router
}

// NewTaskController - new task controller.
func NewTaskController(router *web.Router) {
	controller := TaskController{router: router}
	controller.init()
}

func (s *TaskController) init() {
	taskHandler := handlers.NewTaskHandler()
	// Retrieve task details
	s.router.GET("/api/v1/task/{id}", taskHandler.Get)
	// Update task details
	s.router.PATCH("/api/v1/task/{id}", taskHandler.Patch)
	// Delete task
	s.router.DELETE("/api/v1/task/{id}", taskHandler.Delete)

	statusHandler := handlers.NewStatusHandler()
	// Update task status
	s.router.PATCH("/api/v1/task/{id}/status", statusHandler.Patch)

	badgeHandler := handlers.NewBadgeHandler()
	// Add badge to task
	s.router.POST("/api/v1/task/{id}/badges", badgeHandler.Post)
	// Remove badge from task
	s.router.DELETE("/api/v1/task/{id}/badges/{badgeId}", badgeHandler.Delete)
}
