package app

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/shaninalex/jajirra/internal/domain"
)

func NewAuthController(router *fiber.App) {
	controller := AuthController{router: router}
	controller.setRoutes()
}

type AuthController struct {
	router  *fiber.App
	authApi IAuthApi
}

func (s *AuthController) setRoutes() {
	s.router.Post("/api/auth/hook/register", s.handleHookRegister)
	s.router.Post("/api/auth/hook/verify", s.handleHookVerify)
	s.router.Post("/api/auth/hook/login", s.handleHookLogin)
}

func (s *AuthController) handleHookRegister(ctx *fiber.Ctx) error {
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

func (s *AuthController) handleHookVerify(ctx *fiber.Ctx) error {
	return nil
}

func (s *AuthController) handleHookLogin(ctx *fiber.Ctx) error {
	return nil
}
