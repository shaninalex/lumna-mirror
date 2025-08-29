package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type TaskStatus struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	ProjectID uuid.UUID
	Project   *Project `gorm:"foreignKey:ProjectID;references:ID"`

	Tasks []*Task `gorm:"foreignKey:TaskStatusID;references:ID"`

	Title       string `gorm:"uniqueIndex"`
	Description string
	Complete    bool
	Index       uint
	Config      string
}

func (s *TaskStatus) GetID() uuid.UUID   { return s.ID }
func (s *TaskStatus) SetID(id uuid.UUID) { s.ID = id }
func (s *TaskStatus) SaveConfig(cnf TaskStatusConfig) {
	b, err := json.Marshal(cnf)
	if err != nil {
		panic(err)
	}
	s.Config = string(b)
}
func (s *TaskStatus) GetConfig() *TaskStatusConfig {
	if s.Config == "" {
		return NewTaskStatusConfig()
	}
	var config TaskStatusConfig
	err := json.Unmarshal([]byte(s.Config), &config)
	if err != nil {
		return NewTaskStatusConfig()
	}
	return &config
}

type TaskStatusConfig struct {
	Color string `json:"color,omitempty"`
}

func NewTaskStatusConfig() *TaskStatusConfig {
	return &TaskStatusConfig{
		Color: "default",
	}
}
