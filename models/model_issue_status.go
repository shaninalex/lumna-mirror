package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type IssueStatus struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`

	ProjectID uuid.UUID
	Project   *Project `gorm:"foreignKey:ProjectID;references:ID"`

	Issues []*Issue `gorm:"foreignKey:IssueStatusID;references:ID"`

	Title       string `gorm:"uniqueIndex"`
	Description string
	Complete    bool
	Index       uint
	Config      string
}

func (s *IssueStatus) GetID() uuid.UUID   { return s.ID }
func (s *IssueStatus) SetID(id uuid.UUID) { s.ID = id }
func (s *IssueStatus) SaveConfig(cnf IssueStatusConfig) {
	b, err := json.Marshal(cnf)
	if err != nil {
		panic(err)
	}
	s.Config = string(b)
}
func (s *IssueStatus) GetConfig() *IssueStatusConfig {
	if s.Config == "" {
		return NewIssueStatusConfig()
	}
	var config IssueStatusConfig
	err := json.Unmarshal([]byte(s.Config), &config)
	if err != nil {
		return NewIssueStatusConfig()
	}
	return &config
}

type IssueStatusConfig struct {
	Color string `json:"color,omitempty"`
}

func NewIssueStatusConfig() *IssueStatusConfig {
	return &IssueStatusConfig{
		Color: "default",
	}
}
