package logger

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type ActivityLogService interface {
	GetLog(ctx context.Context, entityType string, entityId uint) ([]models.ActivityLog, error)
}

type activityLogService struct {
	repository repositories.ActivityLogRepository
}

func NewActivityLogService(repository repositories.ActivityLogRepository) ActivityLogService {
	return &activityLogService{
		repository: repository,
	}
}

func (s *activityLogService) GetLog(ctx context.Context, entityType string, entityId uint) ([]models.ActivityLog, error) {
	return s.repository.ListByEntity(ctx, entityType, entityId)
}
