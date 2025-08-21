package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`

	UserID uuid.UUID // Owner
	User   User

	OrganizationID uuid.UUID
	Organization   Organization

	Title string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Issues []Issue `gorm:"foreignKey:ProjectID"`
}

// Implement IObject interface

func (s *Project) GetID() uuid.UUID   { return s.ID }
func (s *Project) SetID(id uuid.UUID) { s.ID = id }

type ProjectRepository struct {
	Repository[*Project]
}

func NewProjectRepository() *ProjectRepository {
	s := &ProjectRepository{}
	return s
}

// Project builder pattern code
type ProjectBuilder struct {
	project *Project
}

func NewProjectBuilder() *ProjectBuilder {
	project := &Project{}
	b := &ProjectBuilder{project: project}
	return b
}

func (b *ProjectBuilder) ID(iD uuid.UUID) *ProjectBuilder {
	b.project.ID = iD
	return b
}

func (b *ProjectBuilder) UserID(userID uuid.UUID) *ProjectBuilder {
	b.project.UserID = userID
	return b
}

func (b *ProjectBuilder) User(user User) *ProjectBuilder {
	b.project.User = user
	return b
}

func (b *ProjectBuilder) OrganizationID(organizationID uuid.UUID) *ProjectBuilder {
	b.project.OrganizationID = organizationID
	return b
}

func (b *ProjectBuilder) Organization(organization Organization) *ProjectBuilder {
	b.project.Organization = organization
	return b
}

func (b *ProjectBuilder) Title(title string) *ProjectBuilder {
	b.project.Title = title
	return b
}

func (b *ProjectBuilder) CreatedAt(createdAt time.Time) *ProjectBuilder {
	b.project.CreatedAt = createdAt
	return b
}

func (b *ProjectBuilder) UpdatedAt(updatedAt time.Time) *ProjectBuilder {
	b.project.UpdatedAt = updatedAt
	return b
}

func (b *ProjectBuilder) Issues(issues []Issue) *ProjectBuilder {
	b.project.Issues = issues
	return b
}

func (b *ProjectBuilder) Build() *Project {
	return b.project
}
