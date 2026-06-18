package email

import "go.uber.org/dig"

func Module(c *dig.Container) error {
	if err := c.Provide(ProvideEmailService); err != nil {
		return err
	}
	if err := c.Provide(NewEmailQueue); err != nil {
		return err
	}

	return nil
}
