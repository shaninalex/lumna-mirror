// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/user/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/user/domain"
)

type UserController struct {
	router *fiber.App
}

// NewUserController - new user controller.
func NewUserController(router *fiber.App) *UserController {
	c := &UserController{
		router: router,
	}
	c.init()
	return c
}

func (s *UserController) init() {
	s.setRoutes()
}

func (s *UserController) setRoutes() {
	h := handler.NewUserHandler(domain.NewUserService())
	s.router.Get("/api/user", h.HandleGetUser)
	s.router.Post("/api/user/settings", h.HandleUpdateSettings)
}
