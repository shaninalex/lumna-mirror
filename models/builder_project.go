package models

import (
	"time"

	"github.com/google/uuid"
)

// ProjectBuilder builder pattern code
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
	b.project.User = &user
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

func (b *ProjectBuilder) Build() *Project {
	return b.project
}
