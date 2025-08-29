package builders

import (
	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/models"
)

// TaskStatusBuilder builder pattern code
type TaskStatusBuilder struct {
	issueStatus *models.TaskStatus
}

func NewIssueStatusBuilder() *TaskStatusBuilder {
	issueStatus := &models.TaskStatus{}
	b := &TaskStatusBuilder{issueStatus: issueStatus}
	return b
}

func (b *TaskStatusBuilder) ID(iD uuid.UUID) *TaskStatusBuilder {
	b.issueStatus.ID = iD
	return b
}

func (b *TaskStatusBuilder) ProjectID(projectID uuid.UUID) *TaskStatusBuilder {
	b.issueStatus.ProjectID = projectID
	return b
}

func (b *TaskStatusBuilder) Project(project *models.Project) *TaskStatusBuilder {
	b.issueStatus.Project = project
	return b
}

func (b *TaskStatusBuilder) Title(title string) *TaskStatusBuilder {
	b.issueStatus.Title = title
	return b
}

func (b *TaskStatusBuilder) Description(description string) *TaskStatusBuilder {
	b.issueStatus.Description = description
	return b
}

func (b *TaskStatusBuilder) Complete(complete bool) *TaskStatusBuilder {
	b.issueStatus.Complete = complete
	return b
}

func (b *TaskStatusBuilder) Index(index uint) *TaskStatusBuilder {
	b.issueStatus.Index = index
	return b
}

func (b *TaskStatusBuilder) Config(config string) *TaskStatusBuilder {
	b.issueStatus.Config = config
	return b
}

func (b *TaskStatusBuilder) Build() *models.TaskStatus {
	return b.issueStatus
}
