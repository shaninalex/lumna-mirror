// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

type AuthFormsHandler struct {
	kratosService kratos.IKratos
	config        base.IConfig
}

func NewAuthFormsHandler(config base.IConfig, kratosService kratos.IKratos) *AuthFormsHandler {
	return &AuthFormsHandler{
		config:        config,
		kratosService: kratosService,
	}
}

func (s *AuthFormsHandler) HandleFormLogin(ctx *fiber.Ctx) error {
	flowId := ctx.Query("flow")
	if flowId == "" {
		return ctx.Redirect(web.GetKratosRedirectUrl(s.config, "/self-service/login/browser"), http.StatusMovedPermanently)
	}

	flow, resp, err := s.kratosService.GetLoginFlow(ctx.Context(), ctx.Get("Cookie"), flowId)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJson(ctx, code, nil, err.Error())
	}
	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}

	return web.Success(ctx, data)
}

func (s *AuthFormsHandler) HandleFormRegister(ctx *fiber.Ctx) error {
	flowId := ctx.Query("flow")
	if flowId == "" {
		return ctx.Redirect(web.GetKratosRedirectUrl(s.config, "/self-service/registration/browser"), http.StatusMovedPermanently)
	}

	flow, resp, err := s.kratosService.GetRegistrationFlow(ctx.Context(), ctx.Get("Cookie"), flowId)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJson(ctx, code, nil, err.Error())
	}
	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}

	return web.Success(ctx, data)
}

func (s *AuthFormsHandler) HandleFormError(ctx *fiber.Ctx) error {
	errorId := ctx.Query("id")
	if errorId == "" {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, "flow id does not provided")
	}

	flow, resp, err := s.kratosService.GetErrorFlow(ctx.Context(), errorId)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJson(ctx, code, nil, err.Error())
	}

	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}
	return web.Success(ctx, data)
}

func (s *AuthFormsHandler) HandleFormVerification(ctx *fiber.Ctx) error {
	flowId := ctx.Query("flow")
	if flowId == "" {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, "flow id does not provided")
	}

	flow, resp, err := s.kratosService.GetVerificationFlow(ctx.Context(), ctx.Get("Cookie"), flowId)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJson(ctx, code, nil, err.Error())
	}

	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}

	return web.Success(ctx, data)
}

func (s *AuthFormsHandler) HandleFormRecovery(ctx *fiber.Ctx) error {
	flowId := ctx.Query("flow")
	if flowId == "" {
		return ctx.Redirect(web.GetKratosRedirectUrl(s.config, "/self-service/recovery/browser"), http.StatusMovedPermanently)
	}

	flow, resp, err := s.kratosService.GetRecoveryFlow(ctx.Context(), ctx.Get("Cookie"), flowId)
	if err != nil {
		code := http.StatusBadRequest
		if resp != nil {
			code = resp.StatusCode
		}
		return web.ReturnJson(ctx, code, nil, err.Error())
	}

	defer resp.Body.Close() //nolint:all

	data, err := flow.ToMap()
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}

	return web.Success(ctx, data)
}
