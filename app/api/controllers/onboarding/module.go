package onboarding

import (
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(services.NewWorkspaceService)
	_ = c.Provide(services.ProvideInvitationService)

	_ = c.Provide(NewOnboardingController)

	return nil
}
