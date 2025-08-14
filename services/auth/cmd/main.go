package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/db"
)

// /api/auth/register
// /api/auth/verify
// /api/auth/login
// /api/auth/restore
// /api/auth/session

func main() {
	args := os.Args
	var configPath string
	if len(args) < 2 || os.Getenv("CONFIG_PATH") != "" {
		configPath = os.Getenv("CONFIG_PATH")
	} else {
		configPath = args[1]
	}

	config := base.NewConfig(configPath)
	database := db.InitDB(config.String("app.dsn"))

	app := fiber.New(fiber.Config{
		AppName: "auth",
	})
	app.Use(logger.New())
	app.Use(csrf.New())
	app.Use(db.NewDBMiddleware(database))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
