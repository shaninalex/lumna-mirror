// Copyright © 2025 Lumna. All rights reserved.

package builders

import (
	"time"

	"gitlab.com/shaninalex/flowreon/internal/db"
)

// ProjectBuilder builder pattern code
type ProjectBuilder struct {
	project *Project
}

// NewProjectBuilder - new project builder.
func NewProjectBuilder() *ProjectBuilder {
	project := &Project{}
	b := &ProjectBuilder{project: project}
	return b
}

// ID - id.
func (b *ProjectBuilder) ID(iD uint) *ProjectBuilder {
	b.project.ID = iD
	return b
}

// UserID - user id.
func (b *ProjectBuilder) UserID(userID uint) *ProjectBuilder {
	b.project.UserID = userID
	return b
}

// Code - project code id.
func (b *ProjectBuilder) Code(code string) *ProjectBuilder {
	b.project.Code = code
	return b
}

// Title - title.
func (b *ProjectBuilder) Title(title string) *ProjectBuilder {
	b.project.Title = title
	return b
}

// CreatedAt - created at.
func (b *ProjectBuilder) CreatedAt(createdAt time.Time) *ProjectBuilder {
	b.project.CreatedAt = createdAt
	return b
}

// UpdatedAt - updated at.
func (b *ProjectBuilder) UpdatedAt(updatedAt time.Time) *ProjectBuilder {
	b.project.UpdatedAt = updatedAt
	return b
}

// Build - builds the value.
func (b *ProjectBuilder) Build() *Project {
	return b.project
}
