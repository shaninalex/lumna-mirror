package builders

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/models"
)

// IssueBuilder builder pattern code
type IssueBuilder struct {
	issue *models.Issue
}

func NewIssueBuilder() *IssueBuilder {
	issue := &models.Issue{}
	b := &IssueBuilder{issue: issue}
	return b
}

func (b *IssueBuilder) ID(iD uuid.UUID) *IssueBuilder {
	b.issue.ID = iD
	return b
}

func (b *IssueBuilder) UserID(userID uuid.UUID) *IssueBuilder {
	b.issue.UserID = userID
	return b
}

func (b *IssueBuilder) User(user models.User) *IssueBuilder {
	b.issue.User = &user
	return b
}

func (b *IssueBuilder) EpicID(epicID uuid.UUID) *IssueBuilder {
	b.issue.EpicID = &epicID
	return b
}

func (b *IssueBuilder) Epic(epic models.Epic) *IssueBuilder {
	b.issue.Epic = &epic
	return b
}

func (b *IssueBuilder) OrganizationID(organizationID uuid.UUID) *IssueBuilder {
	b.issue.OrganizationID = organizationID
	return b
}

func (b *IssueBuilder) Organization(organization models.Organization) *IssueBuilder {
	b.issue.Organization = &organization
	return b
}

func (b *IssueBuilder) SprintID(sprintID uuid.UUID) *IssueBuilder {
	b.issue.SprintID = &sprintID
	return b
}

func (b *IssueBuilder) Sprint(sprint models.Sprint) *IssueBuilder {
	b.issue.Sprint = &sprint
	return b
}

func (b *IssueBuilder) ProjectID(projectID uuid.UUID) *IssueBuilder {
	b.issue.ProjectID = projectID
	return b
}

func (b *IssueBuilder) Project(project models.Project) *IssueBuilder {
	b.issue.Project = &project
	return b
}

func (b *IssueBuilder) Assignee(assignee string) *IssueBuilder {
	b.issue.Assignee = assignee
	return b
}

func (b *IssueBuilder) Type(t models.IssueType) *IssueBuilder {
	b.issue.Type = t
	return b
}

func (b *IssueBuilder) Title(title string) *IssueBuilder {
	b.issue.Title = title
	return b
}

func (b *IssueBuilder) Description(description string) *IssueBuilder {
	b.issue.Description = description
	return b
}

func (b *IssueBuilder) CreatedAt(createdAt time.Time) *IssueBuilder {
	b.issue.CreatedAt = createdAt
	return b
}

func (b *IssueBuilder) UpdatedAt(updatedAt time.Time) *IssueBuilder {
	b.issue.UpdatedAt = updatedAt
	return b
}

func (b *IssueBuilder) Build() *models.Issue {
	return b.issue
}
