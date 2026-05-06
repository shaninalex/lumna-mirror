package repositories

import (
	"context"
	"errors"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gorm.io/gorm"
)

type CredentialRepository interface {
	GetByEmail(ctx context.Context, email string) (*models.Credential, error)
}

type GormCredentialRepository struct {
	db *gorm.DB
}

func NewGormCredentialRepository(db *gorm.DB) CredentialRepository {
	return &GormCredentialRepository{
		db: db,
	}
}

var UserNotFoundError = errors.New("user not found")

func (s *GormCredentialRepository) GetByEmail(ctx context.Context, email string) (*models.Credential, error) {
	credentials := models.Credential{}
	if result := s.db.WithContext(ctx).
		Where(&models.Credential{Email: utils.Pointer(email)}).
		First(&credentials); result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, UserNotFoundError
		}
		return nil, result.Error
	}
	return &credentials, nil
}
