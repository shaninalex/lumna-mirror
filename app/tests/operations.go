package tests

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

func CreateProjectWithName(ctx context.Context, name string) *models.Project {
	repo := repositories.NewProjectRespository()
	project := models.Project{Name: name}
	_ = repo.Create(ctx, &project)
	return &project
}

func CreateProject(ctx context.Context) *models.Project {
	repo := repositories.NewProjectRespository()
	project := models.Project{Name: "test"}
	_ = repo.Create(ctx, &project)
	return &project
}

func CreateBoard(ctx context.Context, projectId uint, name string) *models.Board {
	entry := models.Board{
		Name:      name,
		ProjectId: projectId,
	}
	repo := repositories.NewBoardRepository()
	_ = repo.Create(ctx, &entry)
	return &entry
}

func CreateBoardList(ctx context.Context, boardId uint, name string) *models.BoardList {
	entry := models.BoardList{
		Name:    name,
		BoardId: boardId,
	}
	repo := repositories.NewBoardListRepository()
	_ = repo.Create(ctx, &entry)
	return &entry
}
