// Copyright © 2025 Lumna. All rights reserved.

package builders

import (
	"github.com/shaninalex/lumna/internal/db"
)

// TaskStatusBuilder builder pattern code
type TaskStatusBuilder struct {
	issueStatus *TaskStatus
}

// NewIssueStatusBuilder - new issue status builder.
func NewIssueStatusBuilder() *TaskStatusBuilder {
	issueStatus := &TaskStatus{}
	b := &TaskStatusBuilder{issueStatus: issueStatus}
	return b
}

// ID - id.
func (b *TaskStatusBuilder) ID(iD uint) *TaskStatusBuilder {
	b.issueStatus.ID = iD
	return b
}

// ProjectID - project id.
func (b *TaskStatusBuilder) ProjectID(projectID uint) *TaskStatusBuilder {
	b.issueStatus.ProjectID = projectID
	return b
}

// Title - title.
func (b *TaskStatusBuilder) Title(title string) *TaskStatusBuilder {
	b.issueStatus.Title = title
	return b
}

// Complete - complete.
func (b *TaskStatusBuilder) Complete(complete bool) *TaskStatusBuilder {
	b.issueStatus.Completed = complete
	return b
}

// Index - index.
func (b *TaskStatusBuilder) Index(index uint) *TaskStatusBuilder {
	b.issueStatus.ListIndex = index
	return b
}

// Config - config.
func (b *TaskStatusBuilder) Config(config string) *TaskStatusBuilder {
	b.issueStatus.Config = &config
	return b
}

// Build - builds the value.
func (b *TaskStatusBuilder) Build() *TaskStatus {
	return b.issueStatus
}
