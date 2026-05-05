package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type ActivityLogService struct {
	repository repositories.ActivityLogRepository
}

func NewActivityLogService(repository repositories.ActivityLogRepository) *ActivityLogService {
	return &ActivityLogService{
		repository: repository,
	}
}

func (s *ActivityLogService) GetLog(ctx context.Context, entityType string, entityId uint) ([]models.ActivityLog, error) {
	return s.repository.ListByEntity(ctx, entityType, entityId)
}
