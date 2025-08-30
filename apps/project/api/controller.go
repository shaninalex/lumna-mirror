// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/project/api/handler"
)

func NewProjectController(router *fiber.App) {
	controller := ProjectController{router: router}
	controller.init()
}

type ProjectController struct {
	router *fiber.App
}

func (s *ProjectController) init() {
	h := handler.NewProjectHandler()
	s.router.Get("/api/project/list", h.HandleProjectsList)
	s.router.Get("/api/project/:projectCode/tasks", h.HandleProjectTasksList)
	s.router.Get("/api/project/:projectCode/tasks/:taskCode", h.HandleTaskDetail)
	s.router.Patch("/api/project/:projectCode/tasks/:taskCode", h.HandleTaskUpdate)
	s.router.Patch("/api/project/:projectCode/tasks/:taskCode/status", h.HandleTaskPatchStatus)
}
