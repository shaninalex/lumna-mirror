package domain

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

// ProjectReader - project reader.
type ProjectReader interface {
	List(ctx context.Context) ([]*Project, error)
	GetProject(ctx context.Context, id int64) (*Project, error)
}

// ProjectWriter - project writer
type ProjectWriter interface {
	CreateProject(ctx context.Context, project *Project) (*Project, error)
	UpdateProject(ctx context.Context, project *Project) (*Project, error)
	DeleteProject(ctx context.Context, id int64) error
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
	projects, err := ProjectList(ctx, connection)
	if err != nil {
		return nil, err
	}
	output := make([]*Project, len(projects))
	for i, project := range projects {
		output[i] = project
	}

	return output, nil
}

func (p ProjectService) GetProject(ctx context.Context, id int64) (*Project, error) {
	connection := db.GetDb(ctx)
	project, err := ProjectGetByID(ctx, connection, id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p ProjectService) CreateProject(ctx context.Context, project *Project) (*Project, error) {
	_project := &Project{
		Title: project.Title,
		Code:  utils.GenerateEntityCode("project"),
	}
	err := ProjectSave(ctx, db.GetDb(ctx), _project)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	project.Id = _project.Id
	project.Code = _project.Code
	project.CreatedAt = now
	project.UpdatedAt = now
	return project, nil
}

func (p ProjectService) UpdateProject(ctx context.Context, project *Project) (*Project, error) {
	err := ProjectUpdate(ctx, db.GetDb(ctx), &Project{Id: project.Id, Title: project.Title})
	if err != nil {
		return nil, err
	}
	project, err = p.GetProject(ctx, project.Id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (p ProjectService) DeleteProject(ctx context.Context, id int64) error {
	return ProjectDelete(ctx, db.GetDb(ctx), id)
}
