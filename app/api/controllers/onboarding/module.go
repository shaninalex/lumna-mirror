package onboarding

import (
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(NewOnboardingController)

	return nil
}
