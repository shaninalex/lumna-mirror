// Copyright © 2025 Lumna. All rights reserved.

package domain

import (
	"context"
	"time"

	"github.com/shaninalex/lumna/internal/db"
	"github.com/shaninalex/lumna/internal/utils"
)

type Project struct {
	ID        uint
	Title     string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func MakeProject(project *db.Project) *Project {
	return &Project{
		ID:        project.ID,
		Title:     project.Title,
		Code:      project.Code,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
	}
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
		output[i] = MakeProject(project)
	}

	return output, nil
}

func (p ProjectService) GetProject(ctx context.Context, id uint) (*Project, error) {
	connection := db.GetDb(ctx)
	project, err := db.ProjectGetByID(ctx, connection, id)
	if err != nil {
		return nil, err
	}
	return MakeProject(project), nil
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
	project.Code = _project.Code
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
