// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package builders

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

// OrganizationBuilder builder pattern code
type OrganizationBuilder struct {
	organization *models.Organization
}

func NewOrganizationBuilder() *OrganizationBuilder {
	organization := &models.Organization{}
	b := &OrganizationBuilder{organization: organization}
	return b
}

// ID - id.
func (b *OrganizationBuilder) ID(iD uuid.UUID) *OrganizationBuilder {
	b.organization.ID = iD
	return b
}

// UserID - user id.
func (b *OrganizationBuilder) UserID(userID uuid.UUID) *OrganizationBuilder {
	b.organization.UserID = userID
	return b
}

// User - user.
func (b *OrganizationBuilder) User(user models.User) *OrganizationBuilder {
	b.organization.User = &user
	return b
}

// Title - title.
func (b *OrganizationBuilder) Title(title string) *OrganizationBuilder {
	b.organization.Title = title
	return b
}

// Description - description.
func (b *OrganizationBuilder) Description(description string) *OrganizationBuilder {
	b.organization.Description = description
	return b
}

// CreatedAt - created at.
func (b *OrganizationBuilder) CreatedAt(createdAt time.Time) *OrganizationBuilder {
	b.organization.CreatedAt = createdAt
	return b
}

// UpdatedAt - updated at.
func (b *OrganizationBuilder) UpdatedAt(updatedAt time.Time) *OrganizationBuilder {
	b.organization.UpdatedAt = updatedAt
	return b
}

// Build - builds the value.
func (b *OrganizationBuilder) Build() *models.Organization {
	return b.organization
}
