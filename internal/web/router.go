package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/web/middlewares"
	"gorm.io/gorm"
)

func DefaultRouter(config base.IConfig, db *gorm.DB, name string) *fiber.App {
	router := fiber.New(fiber.Config{
		AppName: name,
	})
	router.Use(recover.New())
	router.Use(logger.New())
	router.Use(csrf.New())
	router.Use(database.NewDBMiddleware(db))
	router.Use(middlewares.NewCommmonMiddleware())
	return router
}
