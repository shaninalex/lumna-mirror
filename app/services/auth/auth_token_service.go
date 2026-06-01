package auth

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type TokenService interface {
	CreateRefreshToken(ctx context.Context, identityID uint, rt *models.RefreshToken) error
	DeleteRefreshToken(ctx context.Context, identityID uint) error
	GetByHash(ctx context.Context, hash string) (*models.RefreshToken, error)
}

var _ TokenService = (*authTokenService)(nil)

type authTokenService struct {
	repository repositories.RefreshTokenRepository
}

func NewAuthTokenService(repository repositories.RefreshTokenRepository) TokenService {
	return &authTokenService{
		repository: repository,
	}
}

func (s *authTokenService) CreateRefreshToken(ctx context.Context, identityID uint, rt *models.RefreshToken) error {
	if err := s.repository.DeleteByIdentityID(ctx, identityID); err != nil {
		return err
	}

	if err := s.repository.Create(ctx, rt); err != nil {
		return err
	}
	return nil
}

func (s *authTokenService) DeleteRefreshToken(ctx context.Context, identityID uint) error {
	return s.repository.DeleteByIdentityID(ctx, identityID)
}

func (s *authTokenService) GetByHash(ctx context.Context, hash string) (*models.RefreshToken, error) {
	return s.repository.GetByHash(ctx, hash)
}
