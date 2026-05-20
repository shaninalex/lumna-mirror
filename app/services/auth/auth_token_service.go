package auth

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type AuthTokenService struct {
	repository repositories.RefreshTokenRepository
}

func NewAuthTokenService(repository repositories.RefreshTokenRepository) *AuthTokenService {
	return &AuthTokenService{
		repository: repository,
	}
}

func (s *AuthTokenService) CreateRefreshToken(ctx context.Context, identityID uint, rt *models.RefreshToken) error {
	if err := s.repository.DeleteByIdentityID(ctx, identityID); err != nil {
		return err
	}

	if err := s.repository.Create(ctx, rt); err != nil {
		return err
	}
	return nil
}

func (s *AuthTokenService) DeleteRefreshToken(ctx context.Context, identityID uint) error {
	return s.repository.DeleteByIdentityID(ctx, identityID)
}

func (s *AuthTokenService) GetByHash(ctx context.Context, hash string) (*models.RefreshToken, error) {
	return s.repository.GetByHash(ctx, hash)
}
