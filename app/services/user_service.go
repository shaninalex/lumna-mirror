package services

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

func (s *UserService) Identity(ctx context.Context, userID uuid.UUID) (*models.Identity, error) {
	identity := models.Identity{}
	if result := s.db.WithContext(ctx).First(&identity, "id = ?", userID); result.Error != nil {
		return nil, result.Error
	}
	return &identity, nil
}
