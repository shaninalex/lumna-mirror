package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ActivityLogRepository interface {
	List(ctx context.Context) ([]models.ActivityLog, error)
	ListByEntity(ctx context.Context, entityType string, entityId uint) ([]models.ActivityLog, error)
	GetByID(ctx context.Context, id uint) (*models.ActivityLog, error)
	Create(ctx context.Context, log *models.ActivityLog) error
	Update(ctx context.Context, log *models.ActivityLog) error
	Delete(ctx context.Context, id uint) error
}

type GormActivityLogRepository struct {
	db *gorm.DB
}

func (r *GormActivityLogRepository) ListByEntity(ctx context.Context, entityType string, entityId uint) ([]models.ActivityLog, error) {
	var logs []models.ActivityLog
	if result := r.db.WithContext(ctx).
		Where("entity_type = ? AND entity_id = ?", entityType, entityId).
		Find(&logs).
		Order("created_at DESC"); result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

func NewGormActivityLogRepository(db *gorm.DB) *GormActivityLogRepository {
	return &GormActivityLogRepository{db: db}
}

func (r *GormActivityLogRepository) List(ctx context.Context) ([]models.ActivityLog, error) {
	var logs []models.ActivityLog

	err := r.db.WithContext(ctx).
		Order("created_at DESC").
		Find(&logs).
		Error

	if err != nil {
		return nil, err
	}

	return logs, nil
}

func (r *GormActivityLogRepository) GetByID(ctx context.Context, id uint) (*models.ActivityLog, error) {
	var log models.ActivityLog

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&log).
		Error

	if err != nil {
		return nil, err
	}

	return &log, nil
}

func (r *GormActivityLogRepository) Create(ctx context.Context, log *models.ActivityLog) error {
	return r.db.WithContext(ctx).
		Create(log).
		Error
}

func (r *GormActivityLogRepository) Update(ctx context.Context, log *models.ActivityLog) error {
	return r.db.WithContext(ctx).
		Model(&models.ActivityLog{}).
		Where("id = ?", log.ID).
		Save(log).
		Error
}

func (r *GormActivityLogRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.ActivityLog{}).
		Error
}
