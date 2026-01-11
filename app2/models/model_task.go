package models

import "time"

type Task struct {
	Id        uint
	BoardId   uint
	ListId    uint
	Name      string
	Done      bool
	Order     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Task) GetId() uint              { return s.Id }
func (s *Task) SetId(u uint)             { s.Id = u }
func (s *Task) GetCreatedAt() time.Time  { return s.CreatedAt }
func (s *Task) GetUpdatedAt() time.Time  { return s.UpdatedAt }
func (s *Task) SetCreatedAt(v time.Time) { s.CreatedAt = v }
func (s *Task) SetUpdatedAt(v time.Time) { s.UpdatedAt = v }
