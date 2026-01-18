package services

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/lumna/app/internal"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
)

type UserService struct {
}

func (s *UserService) Identity(ctx context.Context) (*models.Identity, error) {
	database := db.GetDB(ctx)
	userID := ctx.Value(internal.ContextUserID).(uuid.UUID)
	identity := models.Identity{}
	if result := database.First(&identity, "id = ?", userID); result.Error != nil {
		return nil, result.Error
	}
	return &identity, nil
}
