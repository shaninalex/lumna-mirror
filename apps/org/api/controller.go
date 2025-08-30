// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/org/api/handler"
)

type OrganizationController struct {
	router *fiber.App
}

func NewOrganizationController(router *fiber.App) *OrganizationController {
	c := &OrganizationController{
		router: router,
	}
	c.init()
	return c
}

func (s *OrganizationController) init() {
	s.setRoutes()
}

func (s *OrganizationController) setRoutes() {
	// get by user
	h := handler.NewOrganizationHandler()
	s.router.Get("/api/org", h.HandleGetByUser)
	//s.router.Post("/api/org", h.HandleCreate)
	//s.router.Patch("/api/org/:orgCode", h.HandlePatch)
}
