package activity

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	if err := c.Provide(services.NewActivityLogService); err != nil {
		panic(err)
	}

	if err := c.Provide(NewActivityController); err != nil {
		panic(err)
	}
	return nil
}
