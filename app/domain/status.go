// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"
	"encoding/json"

	"github.com/shaninalex/lumna/app/internal/db"
)

type Status struct {
	ID        uint
	Title     string
	ListIndex uint
	ProjectId uint
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

type StatusReader interface {
	Get(ctx context.Context, id uint) (*Status, error)
	ProjectStatuses(ctx context.Context, projectId uint) ([]*Status, error)
}

type StatusWriter interface {
	Create(ctx context.Context, projectId uint, title string) (*Status, error)
	Patch(ctx context.Context, data *Status) (*Status, error)
	Delete(ctx context.Context, statusId uint) error
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
		status, err := db.TaskStatusByID(ctx, db.GetDb(ctx), uint(statusId))
		if err != nil {
			return err
		}
		status.ListIndex = uint(idx)
		if err = db.TaskStatusUpdate(ctx, db.GetDb(ctx), status); err != nil {
			return err
		}
	}
	return nil
}

func (s StatusService) Get(ctx context.Context, id uint) (*Status, error) {
	status, err := db.TaskStatusByID(ctx, db.GetDb(ctx), id)
	if err != nil {
		return nil, err
	}
	return &Status{
		ID:        status.ID,
		Title:     status.Title,
		ListIndex: status.ListIndex,
		ProjectId: status.ProjectID,
		Config:    status.Config,
	}, nil
}

func (s StatusService) ProjectStatuses(ctx context.Context, projectId uint) ([]*Status, error) {
	dbStatuses, err := db.TaskStatusListByProject(ctx, db.GetDb(ctx), projectId)
	if err != nil {
		return nil, err
	}
	statuses := make([]*Status, len(dbStatuses))
	for i, status := range dbStatuses {
		statuses[i] = &Status{
			ID:        status.ID,
			Title:     status.Title,
			ListIndex: status.ListIndex,
			ProjectId: status.ProjectID,
			Config:    status.Config,
		}
	}
	return statuses, nil
}

func (s StatusService) Create(ctx context.Context, projectId uint, title string) (*Status, error) {
	dbStatus := &db.TaskStatus{
		ProjectID: projectId,
		Title:     title,
	}
	status, err := db.TaskStatusCreate(ctx, db.GetDb(ctx), dbStatus)
	if err != nil {
		return nil, err
	}
	return &Status{
		ID:        status.ID,
		Title:     status.Title,
		ListIndex: status.ListIndex,
		ProjectId: status.ProjectID,
		Config:    status.Config,
	}, nil
}

func (s StatusService) Patch(ctx context.Context, data *Status) (*Status, error) {
	if err := db.TaskStatusUpdate(ctx, db.GetDb(ctx), &db.TaskStatus{
		ID:        data.ID,
		ProjectID: data.ProjectId,
		Title:     data.Title,
		ListIndex: data.ListIndex,
		Config:    data.Config,
	}); err != nil {
		return nil, err
	}
	return data, nil
}

func (s StatusService) Delete(ctx context.Context, statusId uint) error {
	return db.TaskStatusDelete(ctx, db.GetDb(ctx), statusId)
}

// NewStatusService - new status service
func NewStatusService() *StatusService {
	return &StatusService{}
}
