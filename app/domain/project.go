package domain

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

type Project struct {
	Id        int64
	Title     string
	Code      string
	UserId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetID - returns the id.
func (s *Project) GetID() int64 { return s.Id }

// SetID - sets the id.
func (s *Project) SetID(id int64) { s.Id = id }

// GetOwnerID - returns the owner id.
func (s *Project) GetOwnerID() int64 { return s.UserId }

// IsOwner - checks if it is owner.
func (s *Project) IsOwner(entity AuthUser) bool { return entity.GetID() == s.GetOwnerID() }

// GetCreatedAt - returns the created at.
func (s *Project) GetCreatedAt() time.Time { return s.CreatedAt }

// GetUpdatedAt - returns the updated at.
func (s *Project) GetUpdatedAt() time.Time { return s.UpdatedAt }

// GetCreatedBy - returns the created by.
func (s *Project) GetCreatedBy() int64 { return s.GetOwnerID() }

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
