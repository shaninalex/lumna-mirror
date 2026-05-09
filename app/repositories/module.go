package repositories

import (
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(NewGormActivityLogRepository)
	_ = c.Provide(NewGormBoardRepository)
	_ = c.Provide(NewGormColumnRepository)
	_ = c.Provide(NewGormCredentialRepository)
	_ = c.Provide(NewGormIdentityRepository)
	_ = c.Provide(NewGormInvitationRepository)
	_ = c.Provide(NewGormProjectRepository)
	_ = c.Provide(NewGormRefreshTokenRepository)
	_ = c.Provide(NewGormTaskRepository)
	_ = c.Provide(NewGormWorkspaceRepository)

	return nil
}
