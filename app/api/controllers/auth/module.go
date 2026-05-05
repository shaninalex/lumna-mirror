package auth

import (
	authentication "gitlab.com/shaninalex/lumna/app/pkg/auth"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	if err := c.Provide(authentication.NewEmailAuthProvider); err != nil {
		panic(err)
	}
	if err := c.Provide(repositories.NewGormRefreshTokenRepository); err != nil {
		panic(err)
	}
	if err := c.Provide(services.NewAuthTokenService); err != nil {
		panic(err)
	}

	if err := c.Provide(NewAuthContoller); err != nil {
		panic(err)
	}
	return nil
}
