// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"gitlab.com/shaninalex/flowreon/apps/org/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/org/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// OrganizationController - organization controller.
type OrganizationController struct {
	router *web.Router
}

// NewOrganizationController - new organization controller.
func NewOrganizationController(router *web.Router) *OrganizationController {
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
	h := handler.NewOrganizationHandler(domain.NewOrganizationAPI())
	s.router.GET("/api/org", h.HandleGetByUser)
}
