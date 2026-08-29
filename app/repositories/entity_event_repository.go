package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type EntityEventRepository interface {
	Filter(ctx context.Context, query any, args ...any) ([]models.EntityEvent, error)
	ListByEntity(ctx context.Context, entityId int64, entityType string) ([]models.EntityEvent, error)
	ListByEntityIds(ctx context.Context, entityIds []int64, entityType string) ([]models.EntityEvent, error)
}

type gormEntityEventRepository struct {
	db *gorm.DB
}

var _ EntityEventRepository = (*gormEntityEventRepository)(nil)

func NewGormEntityEventRepository(db *gorm.DB) EntityEventRepository {
	return &gormEntityEventRepository{db: db}
}

// Filter implements [EntityEventRepository]
func (s *gormEntityEventRepository) Filter(ctx context.Context, query any, args ...any) ([]models.EntityEvent, error) {
	return gorm.G[models.EntityEvent](s.db).Where(query, args).Find(ctx)
}

// List implements [EntityEventRepository].
func (s *gormEntityEventRepository) ListByEntity(ctx context.Context, entityId int64, entityType string) ([]models.EntityEvent, error) {
	return gorm.G[models.EntityEvent](s.db).
		Where("entity_id = ?", entityId).
		Where("entity_type = ?", entityType).
		Order("created_at").
		Find(ctx)
}

// List implements [EntityEventRepository].
func (s *gormEntityEventRepository) ListByEntityIds(ctx context.Context, entityIds []int64, entityType string) ([]models.EntityEvent, error) {
	return gorm.G[models.EntityEvent](s.db).
		Where("entity_id IN ?", entityIds).
		Where("entity_type = ?", entityType).
		Order("created_at").
		Find(ctx)
}
