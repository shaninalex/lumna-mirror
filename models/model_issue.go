package models

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

func (s *Issue) GetID() uuid.UUID      { return s.ID }
func (s *Issue) SetID(id uuid.UUID)    { s.ID = id }
func (s *Issue) GetOwner() IUser       { return s.User }
func (s *Issue) GetOwnerID() uuid.UUID { return s.UserID }
