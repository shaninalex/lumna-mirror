package app

import (
	"github.com/gofiber/fiber/v2"
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
	s.router.Get("/api/org", s.handleGetByUser)
	s.router.Post("/api/org", s.handleCreate)
	s.router.Patch("/api/org/:id", s.handlePatch)
}

func (s *OrganizationController) handleGetByUser(ctx *fiber.Ctx) error {

	return nil
}

func (s *OrganizationController) handleCreate(ctx *fiber.Ctx) error {

	return nil
}

func (s *OrganizationController) handlePatch(ctx *fiber.Ctx) error {

	return nil
}
