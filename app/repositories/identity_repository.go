package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type IdentityRepository interface {
	GetIdentityByID(ctx context.Context, userID uint) (*models.Identity, error)
	List(ctx context.Context) ([]*models.Identity, error)
}

type GormIdentityRepository struct {
	db *gorm.DB
}

func NewGormIdentityRepository(db *gorm.DB) IdentityRepository {
	return &GormIdentityRepository{
		db: db,
	}
}

func (r *GormIdentityRepository) GetIdentityByID(ctx context.Context, userID uint) (*models.Identity, error) {
	var identity models.Identity

	err := r.db.WithContext(ctx).
		First(&identity, "id = ?", userID).
		Error

	if err != nil {
		return nil, err
	}

	return &identity, nil
}

func (r *GormIdentityRepository) List(ctx context.Context) ([]*models.Identity, error) {
	var identities []*models.Identity
	err := r.db.WithContext(ctx).
		Find(&identities).
		Error

	if err != nil {
		return nil, err
	}
	return identities, nil
}
