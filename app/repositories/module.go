package repositories

import (
	"go.uber.org/dig"
)

func Module(c *dig.Container) error {
	_ = c.Provide(NewGormCredentialRepository)
	_ = c.Provide(NewGormIdentityRepository)
	_ = c.Provide(NewGormInvitationRepository)
	
	_ = c.Provide(NewGormActivityLogRepository)
	_ = c.Provide(NewGormListRepository)
	_ = c.Provide(NewGormStatusRepository)
	_ = c.Provide(NewGormProjectRepository)
	_ = c.Provide(NewGormRefreshTokenRepository)
	_ = c.Provide(NewGormTaskRepository)
	_ = c.Provide(NewGormWorkspaceRepository)

	return nil
}
