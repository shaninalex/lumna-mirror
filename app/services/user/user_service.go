package user

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type Service interface {
	Identity(ctx context.Context, userID uint) (*models.Identity, error)
	List(ctx context.Context) ([]*models.Identity, error)
}

type service struct {
	repository repositories.IdentityRepository
}

var _ Service = (*service)(nil)

func NewUserService(repository repositories.IdentityRepository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Identity(ctx context.Context, userID uint) (*models.Identity, error) {
	return s.repository.GetIdentityByID(ctx, userID)
}

func (s *service) List(ctx context.Context) ([]*models.Identity, error) {
	return s.repository.List(ctx)
}
