package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type RefreshTokenRepository interface {
	Create(ctx context.Context, rt *models.RefreshToken) error
	DeleteByIdentityID(ctx context.Context, identityID uint) error
	GetByHash(ctx context.Context, hash string) (*models.RefreshToken, error)
}

type GormRefreshTokenRepository struct {
	db *gorm.DB
}

func NewGormRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &GormRefreshTokenRepository{db: db}
}

func (r *GormRefreshTokenRepository) Create(
	ctx context.Context,
	rt *models.RefreshToken,
) error {
	return r.db.WithContext(ctx).Create(rt).Error
}

func (r *GormRefreshTokenRepository) DeleteByIdentityID(
	ctx context.Context,
	identityID uint,
) error {
	return r.db.WithContext(ctx).
		Where("identity_id = ?", identityID).
		Delete(&models.RefreshToken{}).
		Error
}

func (r *GormRefreshTokenRepository) GetByHash(
	ctx context.Context,
	hash string,
) (*models.RefreshToken, error) {
	var rt models.RefreshToken

	err := r.db.WithContext(ctx).
		Where("hash = ?", hash).
		First(&rt).
		Error

	if err != nil {
		return nil, err
	}

	return &rt, nil
}
