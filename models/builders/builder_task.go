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

func (b *TaskBuilder) ID(iD uuid.UUID) *TaskBuilder {
	b.task.ID = iD
	return b
}

func (b *TaskBuilder) UserID(userID uuid.UUID) *TaskBuilder {
	b.task.UserID = userID
	return b
}

func (b *TaskBuilder) User(user models.User) *TaskBuilder {
	b.task.User = &user
	return b
}

func (b *TaskBuilder) EpicID(epicID uuid.UUID) *TaskBuilder {
	b.task.EpicID = &epicID
	return b
}

func (b *TaskBuilder) Epic(epic models.Epic) *TaskBuilder {
	b.task.Epic = &epic
	return b
}

func (b *TaskBuilder) OrganizationID(organizationID uuid.UUID) *TaskBuilder {
	b.task.OrganizationID = organizationID
	return b
}

func (b *TaskBuilder) Organization(organization models.Organization) *TaskBuilder {
	b.task.Organization = &organization
	return b
}

func (b *TaskBuilder) SprintID(sprintID uuid.UUID) *TaskBuilder {
	b.task.SprintID = &sprintID
	return b
}

func (b *TaskBuilder) Sprint(sprint models.Sprint) *TaskBuilder {
	b.task.Sprint = &sprint
	return b
}

func (b *TaskBuilder) ProjectID(projectID uuid.UUID) *TaskBuilder {
	b.task.ProjectID = projectID
	return b
}

func (b *TaskBuilder) Project(project models.Project) *TaskBuilder {
	b.task.Project = &project
	return b
}

func (b *TaskBuilder) Assignee(assignee string) *TaskBuilder {
	b.task.Assignee = assignee
	return b
}

func (b *TaskBuilder) Title(title string) *TaskBuilder {
	b.task.Title = title
	return b
}

func (b *TaskBuilder) Description(description string) *TaskBuilder {
	b.task.Description = description
	return b
}

func (b *TaskBuilder) CreatedAt(createdAt time.Time) *TaskBuilder {
	b.task.CreatedAt = createdAt
	return b
}

func (b *TaskBuilder) UpdatedAt(updatedAt time.Time) *TaskBuilder {
	b.task.UpdatedAt = updatedAt
	return b
}

func (b *TaskBuilder) Build() *models.Task {
	return b.task
}
