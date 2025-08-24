package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

	IssueStatusID uuid.UUID
	IssueStatus   *IssueStatus `gorm:"foreignKey:IssueStatusID;references:ID"`

	// Assignee who is assigned to do that task. NOTE: Should be issue_assignee separate relation.
	Assignee string

	Title       string
	Completed   bool
	Description string

	// ListIndex - where in a list of statuses it's currently in
	ListIndex uint

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (s *Issue) GetID() uuid.UUID             { return s.ID }
func (s *Issue) SetID(id uuid.UUID)           { s.ID = id }
func (s *Issue) GetOwner() AuthUser           { return s.User }
func (s *Issue) GetOwnerID() uuid.UUID        { return s.UserID }
func (s *Issue) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }
func (s *Issue) GetCreatedAt() time.Time      { return s.CreatedAt }
func (s *Issue) GetUpdatedAt() time.Time      { return s.UpdatedAt }
func (s *Issue) GetDeletedAt() *time.Time {
	if s.DeletedAt.Valid {
		return &s.DeletedAt.Time
	}
	return nil
}
func (s *Issue) IsDeleted() bool         { return s.DeletedAt.Valid }
func (s *Issue) GetCreatedBy() uuid.UUID { return s.GetOwnerID() }
