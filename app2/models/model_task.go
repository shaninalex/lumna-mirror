package models

import "time"

type Task struct {
	Id        uint
	Name      string
	ListId    uint
	Done      bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Task) GetId() uint             { return s.Id }
func (s *Task) SetId(u uint)            { s.Id = u }
func (s *Task) GetCreatedAt() time.Time { return s.CreatedAt }
func (s *Task) GetUpdatedAt() time.Time { return s.UpdatedAt }
