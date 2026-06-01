package activity

import (
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/logger"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(repositories.NewGormActivityLogRepository)
	if err := c.Provide(logger.NewActivityLogService); err != nil {
		panic(err)
	}

	if err := c.Provide(NewActivityController); err != nil {
		panic(err)
	}
	return nil
}
