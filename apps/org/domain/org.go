// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/models"
	"gorm.io/gorm"
)

// OrganizationManager - organization manager.
type OrganizationManager interface {
	Get(ctx context.Context, userID uuid.UUID) (*models.Organization, error)
}

// OrganizationAPI - organization api.
type OrganizationAPI struct {
}

// NewOrganizationAPI - new organization api.
func NewOrganizationAPI() *OrganizationAPI {
	return &OrganizationAPI{}
}

// Get - returns the value.
func (s *OrganizationAPI) Get(ctx context.Context, userID uuid.UUID) (*models.Organization, error) {
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
