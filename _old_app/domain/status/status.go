package status

import (
	"encoding/json"
)

type Status struct {
	Id        int64   `db:"id" json:"id"`
	Title     string  `db:"title" json:"title"`
	ListIndex int64   `db:"list_index" json:"list_index"`
	ProjectId int64   `db:"project_id" json:"project_id"`
	Config    *string `db:"config" json:"config"`
}

// GetID - returns the id.
func (s *Status) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *Status) SetID(id int64) { s.Id = id }

// SaveConfig - saves the config.
func (s *Status) SaveConfig(cnf TaskStatusConfig) {
	b, err := json.Marshal(cnf)
	if err != nil {
		panic(err)
	}
	res := string(b)
	s.Config = &res
}

// GetConfig - returns the config.
func (s *Status) GetConfig() *TaskStatusConfig {
	if s.Config == nil {
		return NewTaskStatusConfig()
	}
	var config TaskStatusConfig
	err := json.Unmarshal([]byte(*s.Config), &config)
	if err != nil {
		return NewTaskStatusConfig()
	}
	return &config
}

// TaskStatusConfig - task status config.
type TaskStatusConfig struct {
	Color string `json:"color,omitempty"`
}

// NewTaskStatusConfig - new task status config.
func NewTaskStatusConfig() *TaskStatusConfig {
	return &TaskStatusConfig{
		Color: "default",
	}
}
