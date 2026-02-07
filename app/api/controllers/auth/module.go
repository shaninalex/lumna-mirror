package auth

import (
	authentication "gitlab.com/shaninalex/lumna/app/internal/auth"
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(authentication.NewLocalAuthProvider)

	_ = c.Provide(NewAuthContoller)
	return nil
}
