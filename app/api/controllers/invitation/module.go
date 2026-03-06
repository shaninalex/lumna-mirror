package invitation

import (
	"gitlab.com/shaninalex/lumna/app/pkg/email"
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(email.ProvideEmailService)
	_ = c.Provide(services.ProvideInvitationService)

	if err := c.Provide(NewInvitationController); err != nil {
		return err
	}
	return nil
}
