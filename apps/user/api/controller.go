// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"gitlab.com/shaninalex/flowreon/apps/user/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/user/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

type UserController struct {
	router *web.Router
}

// NewUserController - new user controller.
func NewUserController(router *web.Router) *UserController {
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
	s.router.GET("/api/user/me", h.HandleGetUser)
	s.router.POST("/api/user/settings", h.HandleUpdateSettings)
}
