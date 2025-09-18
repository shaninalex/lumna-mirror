// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"gitlab.com/shaninalex/flowreon/apps/auth/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/auth/domain"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// NewAuthController - new auth controller.
func NewAuthController(config base.IConfig, router *web.Router, kratosService kratos.IKratos) {
	controller := AuthController{
		config:        config,
		router:        router,
		kratosService: kratosService,
	}
	controller.setRoutes()
}

// AuthController - auth controller.
type AuthController struct {
	router        *web.Router
	kratosService kratos.IKratos
	config        base.IConfig
}

func (s *AuthController) setRoutes() {

	hooks := handler.NewAuthHooksHandler(domain.NewAuthHookAPI())
	s.router.Use(NewAuthHooksMiddleware(s.kratosService, s.config).Wrap)
	s.router.POST("/api/auth/hook/registration", hooks.HandleHookRegister)
	s.router.POST("/api/auth/hook/verify", hooks.HandleHookVerify)
	s.router.POST("/api/auth/hook/login", hooks.HandleHookLogin)
}
