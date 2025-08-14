package main

import "github.com/gofiber/fiber/v2"

// /api/auth/register
// /api/auth/verify
// /api/auth/login
// /api/auth/restore
// /api/auth/session

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
