// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

// AuthFormsHandler - auth forms handler.
type AuthFormsHandler struct {
	kratosService kratos.IKratos
	config        base.IConfig
}

// NewAuthFormsHandler - new auth forms handler.
func NewAuthFormsHandler(config base.IConfig, kratosService kratos.IKratos) *AuthFormsHandler {
	return &AuthFormsHandler{
		config:        config,
		kratosService: kratosService,
	}
}

// HandleFormLogin - handle form login.
func (s *AuthFormsHandler) HandleFormLogin(ctx *fiber.Ctx) error {
	flowID := ctx.Query("flow")
	if flowID == "" {
		return ctx.Redirect(web.GetKratosRedirectURL(s.config, "/self-service/login/browser"), http.StatusMovedPermanently)
	}

	flow, resp, err := s.kratosService.GetLoginFlow(ctx.Context(), ctx.Get("Cookie"), flowID)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJSON(ctx, code, nil, err.Error())
	}
	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJSON(ctx, http.StatusBadRequest, nil, err.Error())
	}

	return web.Success(ctx, data)
}

// HandleFormRegister - handle form register.
func (s *AuthFormsHandler) HandleFormRegister(ctx *fiber.Ctx) error {
	flowID := ctx.Query("flow")
	if flowID == "" {
		return ctx.Redirect(web.GetKratosRedirectURL(s.config, "/self-service/registration/browser"), http.StatusMovedPermanently)
	}

	flow, resp, err := s.kratosService.GetRegistrationFlow(ctx.Context(), ctx.Get("Cookie"), flowID)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJSON(ctx, code, nil, err.Error())
	}
	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJSON(ctx, http.StatusBadRequest, nil, err.Error())
	}

	return web.Success(ctx, data)
}

// HandleFormError - handle form error.
func (s *AuthFormsHandler) HandleFormError(ctx *fiber.Ctx) error {
	errorID := ctx.Query("id")
	if errorID == "" {
		return web.ReturnJSON(ctx, http.StatusBadRequest, nil, "flow id does not provided")
	}

	flow, resp, err := s.kratosService.GetErrorFlow(ctx.Context(), errorID)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJSON(ctx, code, nil, err.Error())
	}

	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJSON(ctx, http.StatusBadRequest, nil, err.Error())
	}
	return web.Success(ctx, data)
}

// HandleFormVerification - handle form verification.
func (s *AuthFormsHandler) HandleFormVerification(ctx *fiber.Ctx) error {
	flowID := ctx.Query("flow")
	if flowID == "" {
		return web.ReturnJSON(ctx, http.StatusBadRequest, nil, "flow id does not provided")
	}

	flow, resp, err := s.kratosService.GetVerificationFlow(ctx.Context(), ctx.Get("Cookie"), flowID)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJSON(ctx, code, nil, err.Error())
	}

	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJSON(ctx, http.StatusBadRequest, nil, err.Error())
	}

	return web.Success(ctx, data)
}

// HandleFormRecovery - handle form recovery.
func (s *AuthFormsHandler) HandleFormRecovery(ctx *fiber.Ctx) error {
	flowID := ctx.Query("flow")
	if flowID == "" {
		return ctx.Redirect(web.GetKratosRedirectURL(s.config, "/self-service/recovery/browser"), http.StatusMovedPermanently)
	}

	flow, resp, err := s.kratosService.GetRecoveryFlow(ctx.Context(), ctx.Get("Cookie"), flowID)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJSON(ctx, code, nil, err.Error())
	}

	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJSON(ctx, http.StatusBadRequest, nil, err.Error())
	}

	return web.Success(ctx, data)
}
