package repositories

import (
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(NewGormCredentialRepository)
	_ = c.Provide(NewGormIdentityRepository)
	_ = c.Provide(NewGormInvitationRepository)
	_ = c.Provide(NewGormBoardRepository)
	_ = c.Provide(NewGormColumnRepository)
	_ = c.Provide(NewGormProjectRepository)
	_ = c.Provide(NewGormRefreshTokenRepository)
	_ = c.Provide(NewGormTaskRepository)
	_ = c.Provide(NewGormWorkspaceRepository)
	_ = c.Provide(NewGormEmailRepository)

	_ = c.Provide(NewGormTaskStorageRepository)
	_ = c.Provide(NewGormEntityEventRepository)

	return nil
}
