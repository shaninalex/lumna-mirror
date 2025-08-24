package builders

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/models"
)

// TaskBuilder builder pattern code
type TaskBuilder struct {
	issue *models.Task
}

func NewIssueBuilder() *TaskBuilder {
	issue := &models.Task{}
	b := &TaskBuilder{issue: issue}
	return b
}

func (b *TaskBuilder) ID(iD uuid.UUID) *TaskBuilder {
	b.issue.ID = iD
	return b
}

func (b *TaskBuilder) UserID(userID uuid.UUID) *TaskBuilder {
	b.issue.UserID = userID
	return b
}

func (b *TaskBuilder) User(user models.User) *TaskBuilder {
	b.issue.User = &user
	return b
}

func (b *TaskBuilder) EpicID(epicID uuid.UUID) *TaskBuilder {
	b.issue.EpicID = &epicID
	return b
}

func (b *TaskBuilder) Epic(epic models.Epic) *TaskBuilder {
	b.issue.Epic = &epic
	return b
}

func (b *TaskBuilder) OrganizationID(organizationID uuid.UUID) *TaskBuilder {
	b.issue.OrganizationID = organizationID
	return b
}

func (b *TaskBuilder) Organization(organization models.Organization) *TaskBuilder {
	b.issue.Organization = &organization
	return b
}

func (b *TaskBuilder) SprintID(sprintID uuid.UUID) *TaskBuilder {
	b.issue.SprintID = &sprintID
	return b
}

func (b *TaskBuilder) Sprint(sprint models.Sprint) *TaskBuilder {
	b.issue.Sprint = &sprint
	return b
}

func (b *TaskBuilder) ProjectID(projectID uuid.UUID) *TaskBuilder {
	b.issue.ProjectID = projectID
	return b
}

func (b *TaskBuilder) Project(project models.Project) *TaskBuilder {
	b.issue.Project = &project
	return b
}

func (b *TaskBuilder) Assignee(assignee string) *TaskBuilder {
	b.issue.Assignee = assignee
	return b
}

func (b *TaskBuilder) Title(title string) *TaskBuilder {
	b.issue.Title = title
	return b
}

func (b *TaskBuilder) Description(description string) *TaskBuilder {
	b.issue.Description = description
	return b
}

func (b *TaskBuilder) CreatedAt(createdAt time.Time) *TaskBuilder {
	b.issue.CreatedAt = createdAt
	return b
}

func (b *TaskBuilder) UpdatedAt(updatedAt time.Time) *TaskBuilder {
	b.issue.UpdatedAt = updatedAt
	return b
}

func (b *TaskBuilder) Build() *models.Task {
	return b.issue
}
