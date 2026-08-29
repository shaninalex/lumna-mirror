package services

import (
	"context"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type EntityEventService interface {
	Filter(ctx context.Context, query any, args ...any) ([]models.EntityEvent, error)
	ListByEntity(ctx context.Context, entityId int64, entityType string) ([]models.EntityEvent, error)
	ListByEntityIds(ctx context.Context, entityIds []int64, entityType string) ([]models.EntityEvent, error)
}

type entityEventService struct {
	repository repositories.EntityEventRepository
}

var _ EntityEventService = (*entityEventService)(nil)

func NewEntityEventService(repository repositories.EntityEventRepository) EntityEventService {
	s := &entityEventService{
		repository: repository,
	}
	s.init()
	return s
}

func (e *entityEventService) init() {
	fmt.Println("EntitEvent service started.")
}

// Filter implements [EntityEventService].
func (s *entityEventService) Filter(ctx context.Context, query any, args ...any) ([]models.EntityEvent, error) {
	return s.repository.Filter(ctx, query, args)
}

// ListByEntity implements [EntityEventService].
func (s *entityEventService) ListByEntity(ctx context.Context, entityId int64, entityType string) ([]models.EntityEvent, error) {
	return s.repository.ListByEntity(ctx, entityId, entityType)
}

// ListByEntityIds implements [EntityEventService].
func (s *entityEventService) ListByEntityIds(ctx context.Context, entityIds []int64, entityType string) ([]models.EntityEvent, error) {
	if len(entityIds) == 0 {
		return []models.EntityEvent{}, nil
	}
	return s.repository.ListByEntityIds(ctx, entityIds, entityType)
}
