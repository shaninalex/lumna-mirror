// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"
	"encoding/json"
	"time"

	"gitlab.com/shaninalex/flowreon/internal/db"
	"gitlab.com/shaninalex/flowreon/internal/utils"
)

type Project struct {
	ID        uint
	Title     string
	Code      string
	Statuses  []Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status struct {
	ID        uint
	Title     string
	Idx       uint
	Completed bool
	ListIndex uint
	Config    *string
}

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

func MakeProject(project *db.Project, statuses []*db.TaskStatus) *Project {
	_project := &Project{
		ID:        project.ID,
		Title:     project.Title,
		Code:      project.Code,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
	}
	for _, status := range statuses {
		_project.Statuses = append(_project.Statuses, Status{
			ID:    status.ID,
			Title: status.Title,
			Idx:   status.ListIndex,
		})
	}
	return _project
}

// ProjectReader - project reader.
type ProjectReader interface {
	List(ctx context.Context) ([]*Project, error)
	GetProject(ctx context.Context, id uint) (*Project, error)
}

// ProjectWriter - project writer
type ProjectWriter interface {
	CreateProject(ctx context.Context, project *Project) (*Project, error)
	UpdateProject(ctx context.Context, project *Project) (*Project, error)
	DeleteProject(ctx context.Context, id uint) error
}

type ProjectManager interface {
	ProjectReader
	ProjectWriter
}

type ProjectService struct {
}

func NewProjectService() *ProjectService {
	return &ProjectService{}
}

func (p ProjectService) List(ctx context.Context) ([]*Project, error) {
	connection := db.GetDb(ctx)
	projects, err := db.ProjectList(ctx, connection)
	if err != nil {
		return nil, err
	}
	output := make([]*Project, len(projects))
	for i, project := range projects {
		statuses, err := db.TaskStatusListByProject(ctx, connection, project.ID)
		if err != nil {
			return nil, err
		}
		output[i] = MakeProject(project, statuses)
	}

	return output, nil
}

func (p ProjectService) GetProject(ctx context.Context, id uint) (*Project, error) {
	connection := db.GetDb(ctx)
	project, err := db.ProjectGetByID(ctx, connection, id)
	if err != nil {
		return nil, err
	}
	statuses, err := db.TaskStatusListByProject(ctx, connection, project.ID)
	if err != nil {
		return nil, err
	}
	return MakeProject(project, statuses), nil
}

func (p ProjectService) CreateProject(ctx context.Context, project *Project) (*Project, error) {
	_project := &db.Project{
		Title: project.Title,
		Code:  utils.GenerateEntityCode("project"),
	}
	err := db.ProjectSave(ctx, db.GetDb(ctx), _project)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	project.ID = _project.ID
	project.CreatedAt = now
	project.UpdatedAt = now
	return project, nil
}

func (p ProjectService) UpdateProject(ctx context.Context, project *Project) (*Project, error) {
	err := db.ProjectUpdate(ctx, db.GetDb(ctx), &db.Project{ID: project.ID, Title: project.Title})
	if err != nil {
		return nil, err
	}
	project, err = p.GetProject(ctx, project.ID)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p ProjectService) DeleteProject(ctx context.Context, id uint) error {
	return db.ProjectDelete(ctx, db.GetDb(ctx), id)
}
