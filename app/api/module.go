package api

import (
	"gitlab.com/shaninalex/lumna/app/api/controllers/activity"
	"gitlab.com/shaninalex/lumna/app/api/controllers/auth"
	"gitlab.com/shaninalex/lumna/app/api/controllers/column"
	"gitlab.com/shaninalex/lumna/app/api/controllers/invitation"
	"gitlab.com/shaninalex/lumna/app/api/controllers/list"
	"gitlab.com/shaninalex/lumna/app/api/controllers/onboarding"
	"gitlab.com/shaninalex/lumna/app/api/controllers/project"
	"gitlab.com/shaninalex/lumna/app/api/controllers/task"
	"gitlab.com/shaninalex/lumna/app/api/controllers/user"
	"gitlab.com/shaninalex/lumna/app/api/controllers/workspace"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = auth.Module(c)
	_ = list.Module(c)
	_ = column.Module(c)
	_ = project.Module(c)
	_ = task.Module(c)
	_ = user.Module(c)
	_ = activity.Module(c)
	_ = invitation.Module(c)
	_ = onboarding.Module(c)
	_ = workspace.Module(c)

	_ = c.Provide(NewApi)

	return nil
}
