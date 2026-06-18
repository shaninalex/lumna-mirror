package sprint

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type Service interface {
	GetByID(ctx context.Context, SprintId uint) (*models.Sprint, error)
	List(ctx context.Context, query repositories.SprintListQuery) ([]*models.Sprint, error)
	Create(ctx context.Context, sprint *models.Sprint) error
	Update(ctx context.Context, sprint *models.Sprint) error
}

type service struct {
	repository repositories.SprintRepository
}

func NewService(repository repositories.SprintRepository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetByID(ctx context.Context, SprintId uint) (*models.Sprint, error) {
	return s.repository.GetByID(ctx, SprintId)
}

func (s *service) List(ctx context.Context, query repositories.SprintListQuery) ([]*models.Sprint, error) {
	return s.repository.List(ctx, query)
}

func (s *service) Create(ctx context.Context, sprint *models.Sprint) error {
	return s.repository.Create(ctx, sprint)
}

func (s *service) Update(ctx context.Context, sprint *models.Sprint) error {
	return s.repository.Update(ctx, sprint)
}
