package org

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
)

type OrganizationApi struct {
}

func NewOrganizationApi() *OrganizationApi {
	return &OrganizationApi{}
}

func (s *OrganizationApi) Get(ctx context.Context, userID uuid.UUID) (*database.Organization, error) {
	organization := &database.Organization{
		UserID: userID,
	}
	db := database.GetDB(ctx)
	result := db.First(&organization)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, apperrors.OrgNotFound
		}
		return nil, result.Error
	}
	return organization, nil
}
