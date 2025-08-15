package app

import "github.com/gofiber/fiber/v2"

func NewAuthController(router *fiber.App) {
	controller := AuthController{router: router}
	controller.setRoutes()
}

type AuthController struct {
	router *fiber.App
}

func (s *AuthController) setRoutes() {
	s.router.Post("/api/auth/register", s.handleRegister)
	s.router.Post("/api/auth/verify", s.handleVerify)
	s.router.Post("/api/auth/login", s.handleLogin)
	s.router.Post("/api/auth/restore", s.handleRestore)
	s.router.Get("/api/auth/session", s.handleSession)
}

func (s *AuthController) handleRegister(ctx *fiber.Ctx) error {
	return nil
}

func (s *AuthController) handleVerify(ctx *fiber.Ctx) error {
	return nil
}

func (s *AuthController) handleLogin(ctx *fiber.Ctx) error {
	return nil
}

func (s *AuthController) handleRestore(ctx *fiber.Ctx) error {
	return nil
}

func (s *AuthController) handleSession(ctx *fiber.Ctx) error {
	return ctx.JSON(map[string]any{
		"hello": "world",
	})
}
