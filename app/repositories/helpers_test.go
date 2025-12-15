package repositories_test

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

func createUser(ctx context.Context, repo *repositories.UserRepository, email string) *models.User {
	h, _ := utils.CreatePasswordHash(email)
	user := models.User{
		Email:        email,
		PasswordHash: h,
	}
	_ = repo.Create(ctx, &user)

	return &user
}

func createProject(ctx context.Context) *models.Project {
	repo := repositories.NewProjectRespository()
	project := models.Project{Name: "test"}
	_ = repo.Create(ctx, &project)
	return &project
}

func createBoard(ctx context.Context, projectId uint, name string) *models.Board {
	entry := models.Board{
		Name:      name,
		ProjectId: projectId,
	}
	repo := repositories.NewBoardRepository()
	_ = repo.Create(ctx, &entry)
	return &entry
}

func createBoardList(ctx context.Context, boardId uint, name string) *models.BoardList {
	entry := models.BoardList{
		Name:    name,
		BoardId: boardId,
	}
	repo := repositories.NewBoardListRepository()
	_ = repo.Create(ctx, &entry)
	return &entry
}
