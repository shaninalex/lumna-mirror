package jobs

import (
	"gitlab.com/shaninalex/lumna/app/jobs/email"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Invoke(email.Module(c))

	return nil
}
