package project

import (
	"gitlab.com/shaninalex/lumna/app/internal/logger"
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(services.NewProjectService)
	_ = c.Provide(services.NewBoardService)
	_ = c.Provide(logger.ProvideLogger)

	_ = c.Provide(NewProjectsController)
	return nil
}
