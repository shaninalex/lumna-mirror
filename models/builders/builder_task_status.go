// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package builders

import (
	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

// TaskStatusBuilder builder pattern code
type TaskStatusBuilder struct {
	issueStatus *models.TaskStatus
}

// NewIssueStatusBuilder - new issue status builder.
func NewIssueStatusBuilder() *TaskStatusBuilder {
	issueStatus := &models.TaskStatus{}
	b := &TaskStatusBuilder{issueStatus: issueStatus}
	return b
}

// ID - id.
func (b *TaskStatusBuilder) ID(iD uuid.UUID) *TaskStatusBuilder {
	b.issueStatus.ID = iD
	return b
}

// ProjectID - project id.
func (b *TaskStatusBuilder) ProjectID(projectID uuid.UUID) *TaskStatusBuilder {
	b.issueStatus.ProjectID = projectID
	return b
}

// Title - title.
func (b *TaskStatusBuilder) Title(title string) *TaskStatusBuilder {
	b.issueStatus.Title = title
	return b
}

// Description - description.
func (b *TaskStatusBuilder) Description(description string) *TaskStatusBuilder {
	b.issueStatus.Description = description
	return b
}

// Complete - complete.
func (b *TaskStatusBuilder) Complete(complete bool) *TaskStatusBuilder {
	b.issueStatus.Complete = complete
	return b
}

// Index - index.
func (b *TaskStatusBuilder) Index(index uint) *TaskStatusBuilder {
	b.issueStatus.Index = index
	return b
}

// Config - config.
func (b *TaskStatusBuilder) Config(config string) *TaskStatusBuilder {
	b.issueStatus.Config = config
	return b
}

// Build - builds the value.
func (b *TaskStatusBuilder) Build() *models.TaskStatus {
	return b.issueStatus
}
