package org

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
	"gitlab.com/shaninalex/jajirra/models"
	"gorm.io/gorm"
)

type OrganizationApi struct {
}

func NewOrganizationApi() *OrganizationApi {
	return &OrganizationApi{}
}

func (s *OrganizationApi) Get(ctx context.Context, userID uuid.UUID) (*models.Organization, error) {
	var user models.User
	if err := database.GetDB(ctx).Preload("Organization").First(&user, "id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.OrgNotFound
		}
		return nil, err
	}

	if user.Organization == nil {
		return nil, apperrors.OrgNotFound
	}

	return user.Organization, nil
}
