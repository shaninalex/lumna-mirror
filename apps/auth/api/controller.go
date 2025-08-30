// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/apps/auth/api/handler"
	domain2 "gitlab.com/shaninalex/flowreon/apps/auth/domain"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
)

func NewAuthController(config base.IConfig, router *fiber.App, kratosService kratos.IKratos) {
	controller := AuthController{
		config:        config,
		router:        router,
		kratosService: kratosService,
	}
	controller.setRoutes()
}

type AuthController struct {
	router        *fiber.App
	kratosService kratos.IKratos
	config        base.IConfig
}

func (s *AuthController) setRoutes() {
	forms := handler.NewAuthFormsHandler(s.config, s.kratosService)
	s.router.Get("/api/auth/form/login", forms.HandleFormLogin)
	s.router.Get("/api/auth/form/registration", forms.HandleFormRegister)
	s.router.Get("/api/auth/form/error", forms.HandleFormError)
	s.router.Get("/api/auth/form/verification", forms.HandleFormVerification)
	s.router.Get("/api/auth/form/recovery", forms.HandleFormRecovery)

	hooks := handler.NewAuthHooksHandler(domain2.NewAuthHookApi())
	s.router.Post("/api/auth/hook/registration", hooks.HandleHookRegister)
	s.router.Post("/api/auth/hook/verify", hooks.HandleHookVerify)
	s.router.Post("/api/auth/hook/login", hooks.HandleHookLogin)
}
