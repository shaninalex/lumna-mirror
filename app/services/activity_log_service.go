package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ActivityLogService struct {
	db *gorm.DB
}

func NewActivityLogService(db *gorm.DB) *ActivityLogService {
	return &ActivityLogService{
		db: db,
	}
}

func (s *ActivityLogService) GetLog(ctx context.Context, entityType string, entityId uint) ([]models.ActivityLog, error) {
	var logs []models.ActivityLog
	if result := s.db.WithContext(ctx).
		Where("entity_type = ? AND entity_id = ?", entityType, entityId).
		Find(&logs).
		Order("created_at DESC"); result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}
