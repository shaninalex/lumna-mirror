package board

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(services.NewBoardService)
	_ = c.Provide(services.NewColumnService)
	_ = c.Provide(services.NewTaskService)

	_ = c.Provide(NewBoardController)
	return nil
}
