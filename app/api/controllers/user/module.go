package user

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(services.NewUserService)

	_ = c.Provide(NewUserController)
	return nil
}
