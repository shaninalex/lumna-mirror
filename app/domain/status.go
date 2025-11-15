package domain

import (
	"context"
	"encoding/json"

	"gitlab.com/shaninalex/lumna/app/internal/db"
)

type Status struct {
	ID        int64   `db:"id" json:"id"`
	Title     string  `db:"title" json:"title"`
	ListIndex int64   `db:"list_index" json:"list_index"`
	ProjectId int64   `db:"project_id" json:"project_id"`
	Config    *string `db:"config" json:"config"`
}

// GetID - returns the id.
func (s *Status) GetID() int64 { return s.ID }

// SetID - sets the id.
func (s *Status) SetID(id int64) { s.ID = id }

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

type StatusReader interface {
	Get(ctx context.Context, id int64) (*Status, error)
	ProjectStatuses(ctx context.Context, projectId int64) ([]*Status, error)
}

type StatusWriter interface {
	Create(ctx context.Context, projectId int64, title string) (*Status, error)
	Patch(ctx context.Context, data *Status) (*Status, error)
	Delete(ctx context.Context, statusId int64) error
	SortProjectStatus(ctx context.Context, data map[int64]int64) error
}

type StatusManager interface {
	StatusReader
	StatusWriter
}

type StatusService struct {
}

func (s StatusService) SortProjectStatus(ctx context.Context, data map[int64]int64) error {
	for idx, statusId := range data {
		status, err := TaskStatusByID(ctx, db.GetDb(ctx), statusId)
		if err != nil {
			return err
		}
		status.ListIndex = idx
		if err = TaskStatusUpdate(ctx, db.GetDb(ctx), status); err != nil {
			return err
		}
	}
	return nil
}

func (s StatusService) Get(ctx context.Context, id int64) (*Status, error) {
	status, err := TaskStatusByID(ctx, db.GetDb(ctx), id)
	if err != nil {
		return nil, err
	}
	return &Status{
		ID:        status.ID,
		Title:     status.Title,
		ListIndex: status.ListIndex,
		ProjectId: status.ProjectId,
		Config:    status.Config,
	}, nil
}

func (s StatusService) ProjectStatuses(ctx context.Context, projectId int64) ([]*Status, error) {
	dbStatuses, err := TaskStatusListByProject(ctx, db.GetDb(ctx), projectId)
	if err != nil {
		return nil, err
	}
	statuses := make([]*Status, len(dbStatuses))
	for i, status := range dbStatuses {
		statuses[i] = &Status{
			ID:        status.ID,
			Title:     status.Title,
			ListIndex: status.ListIndex,
			ProjectId: status.ProjectId,
			Config:    status.Config,
		}
	}
	return statuses, nil
}

func (s StatusService) Create(ctx context.Context, projectId int64, title string) (*Status, error) {
	status := &Status{
		ProjectId: projectId,
		Title:     title,
	}
	status, err := TaskStatusCreate(ctx, db.GetDb(ctx), status)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (s StatusService) Patch(ctx context.Context, data *Status) (*Status, error) {
	if err := TaskStatusUpdate(ctx, db.GetDb(ctx), data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s StatusService) Delete(ctx context.Context, statusId int64) error {
	return TaskStatusDelete(ctx, db.GetDb(ctx), statusId)
}

// NewStatusService - new status service
func NewStatusService() *StatusService {
	return &StatusService{}
}
