package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Organization struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	// Creator of an Organization
	UserID uuid.UUID
	User   *User

	Title       string
	Description string

	Users []*User `gorm:"foreignKey:OrganizationID;references:ID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// OrganizationBuilder builder pattern code
type OrganizationBuilder struct {
	organization *Organization
}

func NewOrganizationBuilder() *OrganizationBuilder {
	organization := &Organization{}
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

func (b *OrganizationBuilder) User(user User) *OrganizationBuilder {
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

func (b *OrganizationBuilder) Build() *Organization {
	return b.organization
}
