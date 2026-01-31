package services

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

type UserService struct {
}

func (s *UserService) Identity(ctx context.Context, userID uuid.UUID) (*models.Identity, error) {
	identity := models.Identity{}
	if result := db.GetDB(ctx).First(&identity, "id = ?", userID); result.Error != nil {
		return nil, result.Error
	}
	return &identity, nil
}
