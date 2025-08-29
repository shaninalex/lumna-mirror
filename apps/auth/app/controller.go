package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/domain"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gitlab.com/shaninalex/flowreon/internal/web"
)

func NewAuthController(config base.IConfig, router *fiber.App, authApi IAuthApi, kratosService kratos.IKratos) {
	controller := AuthController{
		config:        config,
		router:        router,
		authApi:       authApi,
		kratosService: kratosService,
	}
	controller.setRoutes()
}

type AuthController struct {
	router        *fiber.App
	authApi       IAuthApi
	kratosService kratos.IKratos
	config        base.IConfig
}

func (s *AuthController) setRoutes() {
	// kratos forms
	s.router.Get("/api/auth/form/login", s.HandleFormLogin)
	s.router.Get("/api/auth/form/registration", s.HandleFormRegister)
	s.router.Get("/api/auth/form/error", s.HandleFormError)
	s.router.Get("/api/auth/form/verification", s.HandleFormVerification)
	s.router.Get("/api/auth/form/recovery", s.HandleFormRecovery)

	// hooks
	// TODO: we can authenticate hooks by "ory_kratos_continuity" cookie
	s.router.Post("/api/auth/hook/registration", s.HandleHookRegister)
	s.router.Post("/api/auth/hook/verify", s.HandleHookVerify)
	s.router.Post("/api/auth/hook/login", s.HandleHookLogin)
}

func (s *AuthController) HandleFormLogin(ctx *fiber.Ctx) error {
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

func (s *AuthController) HandleFormRegister(ctx *fiber.Ctx) error {
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

func (s *AuthController) HandleFormError(ctx *fiber.Ctx) error {
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

func (s *AuthController) HandleFormVerification(ctx *fiber.Ctx) error {
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

func (s *AuthController) HandleFormRecovery(ctx *fiber.Ctx) error {
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

func (s *AuthController) HandleHookRegister(ctx *fiber.Ctx) error {
	var data domain.HooksKratosPayloadDTO
	err := ctx.BodyParser(&data)
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}
	err = s.authApi.HookRegister(ctx.Context(), &data)
	if err != nil {
		return web.ReturnJson(ctx, http.StatusBadRequest, nil, err.Error())
	}
	return web.Success(ctx, nil)
}

func (s *AuthController) HandleHookVerify(ctx *fiber.Ctx) error {
	return web.Success(ctx, nil)
}

func (s *AuthController) HandleHookLogin(ctx *fiber.Ctx) error {
	return web.Success(ctx, nil)
}
