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

func (b *OrganizationBuilder) ID(iD uuid.UUID) *OrganizationBuilder {
	b.organization.ID = iD
	return b
}

func (b *OrganizationBuilder) UserID(userID uuid.UUID) *OrganizationBuilder {
	b.organization.UserID = userID
	return b
}

func (b *OrganizationBuilder) User(user models.User) *OrganizationBuilder {
	b.organization.User = &user
	return b
}

func (b *OrganizationBuilder) Title(title string) *OrganizationBuilder {
	b.organization.Title = title
	return b
}

func (b *OrganizationBuilder) Description(description string) *OrganizationBuilder {
	b.organization.Description = description
	return b
}

func (b *OrganizationBuilder) CreatedAt(createdAt time.Time) *OrganizationBuilder {
	b.organization.CreatedAt = createdAt
	return b
}

func (b *OrganizationBuilder) UpdatedAt(updatedAt time.Time) *OrganizationBuilder {
	b.organization.UpdatedAt = updatedAt
	return b
}

func (b *OrganizationBuilder) Build() *models.Organization {
	return b.organization
}
