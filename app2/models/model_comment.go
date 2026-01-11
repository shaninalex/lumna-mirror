package models

import (
	"time"
)

type Comment struct {
	Id         uint
	Date       time.Time
	Content    string
	EntityType string
	EntityId   uint
	AuthorId   uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (s *Comment) GetId() uint                { return s.Id }
func (s *Comment) SetId(v uint)               { s.Id = v }
func (s *Comment) GetOwnerId() uint           { return s.AuthorId }
func (s *Comment) IsOwner(user AuthUser) bool { return user.GetId() == s.AuthorId }
func (s *Comment) GetCreatedAt() time.Time    { return s.CreatedAt }
func (s *Comment) GetUpdatedAt() time.Time    { return s.UpdatedAt }
func (s *Comment) SetCreatedAt(v time.Time)   { s.CreatedAt = v }
func (s *Comment) SetUpdatedAt(v time.Time)   { s.UpdatedAt = v }
