package api

import (
	"gitlab.com/shaninalex/lumna/app/api/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/api/controllers/board"
	"gitlab.com/shaninalex/lumna/app/api/controllers/column"
	"gitlab.com/shaninalex/lumna/app/api/controllers/project"
	"gitlab.com/shaninalex/lumna/app/api/controllers/task"
	"gitlab.com/shaninalex/lumna/app/api/controllers/user"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = auth.Module(c)
	_ = board.Module(c)
	_ = column.Module(c)
	_ = project.Module(c)
	_ = task.Module(c)
	_ = user.Module(c)

	if err := c.Provide(NewApi); err != nil {
		panic(err)
	}

	return nil
}
