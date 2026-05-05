package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type AuthTokenService struct {
	db *gorm.DB
}

func NewAuthTokenService(db *gorm.DB) *AuthTokenService {
	return &AuthTokenService{
		db: db,
	}
}

func (s *AuthTokenService) CreateRefreshToken(ctx context.Context, identityID uint, rt *models.RefreshToken) error {
	if result := s.db.WithContext(ctx).Where("identity_id = ?", identityID).Delete(&models.RefreshToken{}); result.Error != nil {
		return result.Error
	}

	if result := s.db.WithContext(ctx).Create(&rt); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *AuthTokenService) DeleteRefreshToken(ctx context.Context, identityID uint) error {
	return s.db.WithContext(ctx).Where("identity_id = ?", identityID).Delete(&models.RefreshToken{}).Error
}

func (s *AuthTokenService) GetByHash(ctx context.Context, hash string) (*models.RefreshToken, error) {
	rt := models.RefreshToken{}
	if result := s.db.WithContext(ctx).Preload("Identity").Where("hash = ?", hash).First(&rt); result.Error != nil {
		return nil, result.Error
	}
	return &rt, nil
}
