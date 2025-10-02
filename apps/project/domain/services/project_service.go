// Copyright © 2025 Lumna. All rights reserved.

package services

import (
	"context"
	"time"

	"gitlab.com/shaninalex/flowreon/apps/project/domain/models"
	"gitlab.com/shaninalex/flowreon/internal/db"
)

func MakeProject(project *db.Project, statuses []*db.TaskStatus) *models.Project {
	_project := &models.Project{
		ID:        project.ID,
		Title:     project.Title,
		Code:      project.Code,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
	}
	for _, status := range statuses {
		_project.Statuses = append(_project.Statuses, models.Status{
			ID:    status.ID,
			Title: status.Title,
			Idx:   status.ListIndex,
		})
	}
	return _project
}

// ProjectReader - project reader.
type ProjectReader interface {
	List(ctx context.Context) ([]*models.Project, error)
	GetProject(ctx context.Context, id uint) (*models.Project, error)
}

// ProjectWriter - project writer
type ProjectWriter interface {
	CreateProject(ctx context.Context, project *models.Project) (*models.Project, error)
	UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error)
	DeleteProject(ctx context.Context, id uint) error
}

type ProjectManager interface {
	ProjectReader
	ProjectWriter
}

type ProjectService struct {
}

func (p ProjectService) List(ctx context.Context) ([]*models.Project, error) {
	connection := db.GetDb(ctx)
	projects, err := db.ProjectList(ctx, connection)
	if err != nil {
		return nil, err
	}
	output := make([]*models.Project, len(projects))
	for _, project := range projects {
		statuses, err := db.TaskStatusListByProject(ctx, connection, project.ID)
		if err != nil {
			return nil, err
		}
		output = append(output, MakeProject(project, statuses))
	}

	return output, nil
}

func (p ProjectService) GetProject(ctx context.Context, id uint) (*models.Project, error) {
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

func (p ProjectService) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	_project := &db.Project{
		Title: project.Title,
		Code:  project.Code,
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

func (p ProjectService) UpdateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
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
