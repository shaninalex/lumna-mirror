package models

import "time"

type Sprint struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Title       string     `gorm:"not null" json:"title"`
	Description string     `json:"body"`
	Done        bool       `json:"done"`
	ProjectID   uint       `json:"project_id"`
	StartedAt   *time.Time `json:"started_at"`
	FinishedAt  *time.Time `json:"finished_at"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (s *Sprint) GetTitle() string {
	return s.Title
}
