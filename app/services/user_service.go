package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type UserManager interface {
	Identity(ctx context.Context, userID uint) (*models.Identity, error)
	List(ctx context.Context) ([]*models.Identity, error)
}

type UserService struct {
	repository repositories.IdentityRepository
}

func NewUserService(repository repositories.IdentityRepository) UserManager {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Identity(ctx context.Context, userID uint) (*models.Identity, error) {
	return s.repository.GetIdentityByID(ctx, userID)
}

func (s *UserService) List(ctx context.Context) ([]*models.Identity, error) {
	return s.repository.List(ctx)
}
