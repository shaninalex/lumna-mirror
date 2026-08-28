package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/testutils"
)

func Test_TaskRepository_CreateTask(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	taskRepo := repositories.NewGormTaskRepository(db)

	workspace := &models.Workspace{Title: "ws", Active: true}
	db.WithContext(ctx).Create(workspace)

	project := &models.Project{Title: "project", WorkspaceID: workspace.ID}
	db.WithContext(ctx).Create(project)

	board := &models.Board{Title: "board", ProjectID: project.ID}
	db.WithContext(ctx).Create(board)

	column := &models.Column{Title: "column", BoardId: uint(board.ID)}
	db.WithContext(ctx).Create(column)

	data := &models.TaskCreateOnBoard{
		Title:    "task",
		Body:     "body",
		BoardId:  utils.Pointer(uint(board.ID)),
		ColumnId: utils.Pointer(uint(column.ID)),
	}

	task := models.Task{
		Title: data.Title,
		Body:  data.Body,
	}

	err := taskRepo.Create(ctx, &task)
	assert.NoError(t, err)

	err = taskRepo.CreateTaskBoard(ctx, models.BoardTask{
		TaskId:   uint(task.ID),
		Position: uint(0),
		BoardId:  *data.BoardId,
		ColumnId: *data.ColumnId,
	})
	assert.NoError(t, err)
}
