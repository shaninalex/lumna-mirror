package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/base"
	"gitlab.com/shaninalex/jajirra/internal/kratos"
	"gorm.io/gorm"
)

func DefaultRouter(db *gorm.DB, name string) *fiber.App {
	router := fiber.New(fiber.Config{
		AppName: name,
	})
	router.Use(recover.New())
	router.Use(logger.New())
	router.Use(database.NewDbMiddleware(db))
	router.Use(NewCommonMiddleware())
	return router
}

func AuthRouter(cnf base.IConfig, db *gorm.DB, name string) *fiber.App {
	router := DefaultRouter(db, name)
	kratosClient := kratos.NewKratosService(cnf.String("kratos.url_browser"))
	router.Use(NewAuthMiddleware(kratosClient))
	return router
}
