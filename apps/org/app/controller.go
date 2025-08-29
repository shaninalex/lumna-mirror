package app

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/internal/org"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

type OrganizationController struct {
	router *fiber.App
	api    *org.OrganizationApi
}

func NewOrganizationController(router *fiber.App) *OrganizationController {
	c := &OrganizationController{
		router: router,
		api:    org.NewOrganizationApi(),
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
	organization, err := s.api.Get(ctx.Context(), web.GetUserId(ctx))
	if errors.Is(err, apperrors.OrgNotFound) {
		return web.ReturnJson(ctx, http.StatusNotFound, nil, err.Error())
	}
	return web.Success(ctx, ToDto(organization))
}

func (s *OrganizationController) handleCreate(ctx *fiber.Ctx) error {
	return nil
}

func (s *OrganizationController) handlePatch(ctx *fiber.Ctx) error {
	return nil
}
