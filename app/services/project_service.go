package services

import (
	"context"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type ProjectManager interface {
	List(ctx context.Context) ([]*models.Project, error)
	Create(ctx context.Context, entry *models.Project) error
}

type ProjectService struct {
	projectRepository *repositories.ProjectRespository
	boardRepository   *repositories.BoardRepository
}

func NewProjectManager() ProjectManager {
	return &ProjectService{
		projectRepository: repositories.NewProjectRespository(),
		boardRepository:   repositories.NewBoardRepository(),
	}
}

var _ ProjectManager = (*ProjectService)(nil)

func (s *ProjectService) List(ctx context.Context) ([]*models.Project, error) {
	projects, err := s.projectRepository.List(ctx, nil)
	if err != nil {
		return nil, err
	}

	ids := make([]uint, len(projects))
	for _, p := range projects {
		ids[0] = p.GetId()
	}
	if boards, err := s.boardRepository.List(ctx, db.Eq("project_id", ids)); err == nil {
		for _, b := range boards {
			for _, p := range projects {
				if p.GetId() == b.GetId() {
					p.Boards = append(p.Boards, b)
				}
			}
		}
	}

	return projects, nil
}

func (s *ProjectService) Create(ctx context.Context, entry *models.Project) error {
	if entry.Name == "" {
		return fmt.Errorf("project name is required")
	}
	count, err := s.projectRepository.Count(ctx, db.Eq("name", entry.Name))
	if err != nil {
		return err
	}
	if count > 0 {
		return fmt.Errorf("project with name %s already exist", entry.Name)
	}
	return s.projectRepository.Create(ctx, entry)
}
