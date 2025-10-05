// Copyright © 2025 Lumna. All rights reserved.

package builders

import (
	"time"

	"gitlab.com/shaninalex/flowreon/internal/db"
)

// TaskBuilder builder pattern code
type TaskBuilder struct {
	task *Task
}

// NewTaskBuilder - new issue builder.
func NewTaskBuilder() *TaskBuilder {
	task := &Task{}
	b := &TaskBuilder{task: task}
	return b
}

// ID - id.
func (b *TaskBuilder) ID(iD uint) *TaskBuilder {
	b.task.ID = iD
	return b
}

// UserID - user id.
func (b *TaskBuilder) UserID(userID uint) *TaskBuilder {
	b.task.UserID = userID
	return b
}

// ProjectID - project id.
func (b *TaskBuilder) ProjectID(projectID uint) *TaskBuilder {
	b.task.ProjectID = projectID
	return b
}

// TaskStatusID - task status id.
func (b *TaskBuilder) TaskStatusID(taskStatusID uint) *TaskBuilder {
	b.task.StatusID = taskStatusID
	return b
}

// Code - set task code
func (b *TaskBuilder) Code(code string) *TaskBuilder {
	b.task.Code = code
	return b
}

// Title - title.
func (b *TaskBuilder) Title(title string) *TaskBuilder {
	b.task.Title = title
	return b
}

// Description - description.
func (b *TaskBuilder) Description(description string) *TaskBuilder {
	b.task.Description = &description
	return b
}

// CreatedAt - created at.
func (b *TaskBuilder) CreatedAt(createdAt time.Time) *TaskBuilder {
	b.task.CreatedAt = createdAt
	return b
}

// UpdatedAt - updated at.
func (b *TaskBuilder) UpdatedAt(updatedAt time.Time) *TaskBuilder {
	b.task.UpdatedAt = updatedAt
	return b
}

// Build - builds the value.
func (b *TaskBuilder) Build() *Task {
	return b.task
}
