package sprint

import (
	"gitlab.com/shaninalex/lumna/app/services/sprint"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(sprint.NewService)

	_ = c.Provide(NewSprintController)
	return nil
}
