// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package builders

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

// ProjectBuilder builder pattern code
type ProjectBuilder struct {
	project *models.Project
}

func NewProjectBuilder() *ProjectBuilder {
	project := &models.Project{}
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

func (b *ProjectBuilder) User(user models.User) *ProjectBuilder {
	b.project.User = &user
	return b
}

func (b *ProjectBuilder) OrganizationID(organizationID uuid.UUID) *ProjectBuilder {
	b.project.OrganizationID = organizationID
	return b
}

func (b *ProjectBuilder) Organization(organization models.Organization) *ProjectBuilder {
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

func (b *ProjectBuilder) Build() *models.Project {
	return b.project
}
