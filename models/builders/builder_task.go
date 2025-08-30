// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package builders

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

// TaskBuilder builder pattern code
type TaskBuilder struct {
	task *models.Task
}

func NewIssueBuilder() *TaskBuilder {
	task := &models.Task{}
	b := &TaskBuilder{task: task}
	return b
}

// ID - id.
func (b *TaskBuilder) ID(iD uuid.UUID) *TaskBuilder {
	b.task.ID = iD
	return b
}

// UserID - user id.
func (b *TaskBuilder) UserID(userID uuid.UUID) *TaskBuilder {
	b.task.UserID = userID
	return b
}

// User - user.
func (b *TaskBuilder) User(user models.User) *TaskBuilder {
	b.task.User = &user
	return b
}

// EpicID - epic id.
func (b *TaskBuilder) EpicID(epicID uuid.UUID) *TaskBuilder {
	b.task.EpicID = &epicID
	return b
}

// Epic - epic.
func (b *TaskBuilder) Epic(epic models.Epic) *TaskBuilder {
	b.task.Epic = &epic
	return b
}

// OrganizationID - organization id.
func (b *TaskBuilder) OrganizationID(organizationID uuid.UUID) *TaskBuilder {
	b.task.OrganizationID = organizationID
	return b
}

// Organization - organization.
func (b *TaskBuilder) Organization(organization models.Organization) *TaskBuilder {
	b.task.Organization = &organization
	return b
}

// SprintID - sprint id.
func (b *TaskBuilder) SprintID(sprintID uuid.UUID) *TaskBuilder {
	b.task.SprintID = &sprintID
	return b
}

// Sprint - sprint.
func (b *TaskBuilder) Sprint(sprint models.Sprint) *TaskBuilder {
	b.task.Sprint = &sprint
	return b
}

// ProjectID - project id.
func (b *TaskBuilder) ProjectID(projectID uuid.UUID) *TaskBuilder {
	b.task.ProjectID = projectID
	return b
}

// Project - project.
func (b *TaskBuilder) Project(project models.Project) *TaskBuilder {
	b.task.Project = &project
	return b
}

// Assignee - assignee.
func (b *TaskBuilder) Assignee(assignee string) *TaskBuilder {
	b.task.Assignee = assignee
	return b
}

// Title - title.
func (b *TaskBuilder) Title(title string) *TaskBuilder {
	b.task.Title = title
	return b
}

// Description - description.
func (b *TaskBuilder) Description(description string) *TaskBuilder {
	b.task.Description = description
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
func (b *TaskBuilder) Build() *models.Task {
	return b.task
}
