// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/project/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
)

func NewProjectController(router *fiber.App) {
	controller := ProjectController{router: router}
	controller.init()
}

type ProjectController struct {
	router *fiber.App
}

func (s *ProjectController) init() {
	projectHandler := handler.NewProjectHandler(domain.NewProjectManagement())
	s.router.Get("/api/project/list", projectHandler.HandleProjectsList)
	s.router.Get("/api/project/:projectCode/tasks", projectHandler.HandleProjectTasksList)

	taskHandler := handler.NewTaskHandler(domain.NewProjectManagement())
	s.router.Get("/api/project/:projectCode/tasks/:taskCode", taskHandler.HandleTaskDetail)
	s.router.Patch("/api/project/:projectCode/tasks/:taskCode", taskHandler.HandleTaskUpdate)
	s.router.Patch("/api/project/:projectCode/tasks/:taskCode/status", taskHandler.HandleTaskPatchStatus)
}
