package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueType string

const (
	IssueTypeStory IssueType = "story"
	IssueTypeTask  IssueType = "task"
	IssueTypeBug   IssueType = "bug"
)

type Issue struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	UserID uuid.UUID
	User   *User

	EpicID *uuid.UUID
	Epic   *Epic

	OrganizationID uuid.UUID
	Organization   *Organization

	SprintID *uuid.UUID
	Sprint   *Sprint

	ProjectID uuid.UUID
	Project   *Project

	Assignee string // TODO: issue_assignee separate relation table

	Type        IssueType
	Title       string
	Description string
	Status      string // in_progress, done, todo...
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// IssueBuilder builder pattern code
type IssueBuilder struct {
	issue *Issue
}

func NewIssueBuilder() *IssueBuilder {
	issue := &Issue{}
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

func (b *IssueBuilder) User(user User) *IssueBuilder {
	b.issue.User = &user
	return b
}

func (b *IssueBuilder) EpicID(epicID uuid.UUID) *IssueBuilder {
	b.issue.EpicID = &epicID
	return b
}

func (b *IssueBuilder) Epic(epic Epic) *IssueBuilder {
	b.issue.Epic = &epic
	return b
}

func (b *IssueBuilder) OrganizationID(organizationID uuid.UUID) *IssueBuilder {
	b.issue.OrganizationID = organizationID
	return b
}

func (b *IssueBuilder) Organization(organization Organization) *IssueBuilder {
	b.issue.Organization = &organization
	return b
}

func (b *IssueBuilder) SprintID(sprintID uuid.UUID) *IssueBuilder {
	b.issue.SprintID = &sprintID
	return b
}

func (b *IssueBuilder) Sprint(sprint Sprint) *IssueBuilder {
	b.issue.Sprint = &sprint
	return b
}

func (b *IssueBuilder) ProjectID(projectID uuid.UUID) *IssueBuilder {
	b.issue.ProjectID = projectID
	return b
}

func (b *IssueBuilder) Project(project Project) *IssueBuilder {
	b.issue.Project = &project
	return b
}

func (b *IssueBuilder) Assignee(assignee string) *IssueBuilder {
	b.issue.Assignee = assignee
	return b
}

func (b *IssueBuilder) Type(t IssueType) *IssueBuilder {
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

func (b *IssueBuilder) Status(status string) *IssueBuilder {
	b.issue.Status = status
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

func (b *IssueBuilder) Build() *Issue {
	return b.issue
}
