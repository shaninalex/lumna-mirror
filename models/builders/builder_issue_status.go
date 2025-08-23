package builders

import (
	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/models"
)

// IssueStatusBuilder builder pattern code
type IssueStatusBuilder struct {
	issueStatus *models.IssueStatus
}

func NewIssueStatusBuilder() *IssueStatusBuilder {
	issueStatus := &models.IssueStatus{}
	b := &IssueStatusBuilder{issueStatus: issueStatus}
	return b
}

func (b *IssueStatusBuilder) ID(iD uuid.UUID) *IssueStatusBuilder {
	b.issueStatus.ID = iD
	return b
}

func (b *IssueStatusBuilder) ProjectID(projectID uuid.UUID) *IssueStatusBuilder {
	b.issueStatus.ProjectID = projectID
	return b
}

func (b *IssueStatusBuilder) Project(project *models.Project) *IssueStatusBuilder {
	b.issueStatus.Project = project
	return b
}

func (b *IssueStatusBuilder) Title(title string) *IssueStatusBuilder {
	b.issueStatus.Title = title
	return b
}

func (b *IssueStatusBuilder) Description(description string) *IssueStatusBuilder {
	b.issueStatus.Description = description
	return b
}

func (b *IssueStatusBuilder) Complete(complete bool) *IssueStatusBuilder {
	b.issueStatus.Complete = complete
	return b
}

func (b *IssueStatusBuilder) Index(index uint) *IssueStatusBuilder {
	b.issueStatus.Index = index
	return b
}

func (b *IssueStatusBuilder) Config(config string) *IssueStatusBuilder {
	b.issueStatus.Config = config
	return b
}

func (b *IssueStatusBuilder) Build() *models.IssueStatus {
	return b.issueStatus
}
