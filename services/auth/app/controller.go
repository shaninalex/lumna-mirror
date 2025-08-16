package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/domain"
	"gitlab.com/shaninalex/jajirra/internal/kratos"
	"gitlab.com/shaninalex/jajirra/internal/web"
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
	// hooks
	s.router.Post("/api/auth/hook/register", s.HandleHookRegister)
	s.router.Post("/api/auth/hook/verify", s.HandleHookVerify)
	s.router.Post("/api/auth/hook/login", s.HandleHookLogin)

	// kratos forms
	s.router.Get("/api/auth/form/login", s.HandleFormLogin)
}

func (s *AuthController) HandleHookRegister(ctx *fiber.Ctx) error {
	var data domain.HooksKratosPayloadDTO
	err := ctx.BodyParser(&data)
	if err != nil {
		return err
	}
	err = s.authApi.HookRegister(ctx.Context(), &data)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return err
	}
	ctx.Status(http.StatusOK)
	return nil
}

func (s *AuthController) HandleHookVerify(ctx *fiber.Ctx) error {
	return nil
}

func (s *AuthController) HandleHookLogin(ctx *fiber.Ctx) error {
	return nil
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
