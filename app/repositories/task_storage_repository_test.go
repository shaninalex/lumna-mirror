package repositories_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/repositories/storage"
	"gitlab.com/shaninalex/lumna/testutils"
	"gorm.io/gorm"
)

func Test_TaskStorageRepository_Save(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	repo := repositories.NewGormTaskStorageRepository(db)

	owner := &models.Identity{FullName: "owner", Email: "owner@example.com"}
	require.NoError(t, db.WithContext(ctx).Create(owner).Error)
	assignee := &models.Identity{FullName: "assignee", Email: "assignee@example.com"}
	require.NoError(t, db.WithContext(ctx).Create(assignee).Error)

	workspace := &models.Workspace{Title: "ws", Active: true}
	require.NoError(t, db.WithContext(ctx).Create(workspace).Error)
	project := &models.Project{Title: "project", WorkspaceID: workspace.ID}
	require.NoError(t, db.WithContext(ctx).Create(project).Error)
	board := &models.Board{Title: "board", ProjectID: project.ID}
	require.NoError(t, db.WithContext(ctx).Create(board).Error)
	column := &models.Column{Title: "column", BoardId: uint(board.ID)}
	require.NoError(t, db.WithContext(ctx).Create(column).Error)

	task := &models.Task{
		Title:        "task",
		Body:         "body",
		ProjectId:    int64(project.ID),
		OwnerId:      int64(owner.ID),
		AssigneesIDs: []int64{int64(assignee.ID), int64(assignee.ID)},
		Boards: []models.TaskBoard{
			{BoardId: int64(board.ID), ColumnId: int64(column.ID), Position: 1},
		},
	}

	require.NoError(t, repo.Save(ctx, task))
	assert.NotZero(t, task.ID, "generated id is written back into the model")

	record, err := gorm.G[storage.TaskRecord](db).Where("id = ?", task.ID).First(ctx)
	require.NoError(t, err)
	assert.Equal(t, "task", record.Title)
	assert.Equal(t, "body", record.Body)
	assert.False(t, record.Completed)

	owners, err := gorm.G[storage.TaskOwnerRecord](db).Where("task_id = ?", task.ID).Find(ctx)
	require.NoError(t, err)
	require.Len(t, owners, 1)
	assert.Equal(t, int64(owner.ID), owners[0].UserID)

	assignees, err := gorm.G[storage.TaskAssigneeRecord](db).Where("task_id = ?", task.ID).Find(ctx)
	require.NoError(t, err)
	require.Len(t, assignees, 1, "duplicated assignees are collapsed")

	boards, err := gorm.G[storage.BoardTaskRecord](db).Where("task_id = ?", task.ID).Find(ctx)
	require.NoError(t, err)
	require.Len(t, boards, 1)
	require.NotNil(t, boards[0].ColumnID)
	assert.Equal(t, int64(column.ID), boards[0].ColumnID)
	assert.Equal(t, int64(1), boards[0].Position)

	// Second save updates in place and replaces the related rows.
	task.Title = "renamed"
	task.Completed = true
	task.OwnerId = 0
	task.AssigneesIDs = nil
	task.Boards[0].ColumnId = 0
	task.Boards[0].Position = 5

	require.NoError(t, repo.Save(ctx, task))

	record, err = gorm.G[storage.TaskRecord](db).Where("id = ?", task.ID).First(ctx)
	require.NoError(t, err)
	assert.Equal(t, "renamed", record.Title)
	assert.True(t, record.Completed)

	owners, err = gorm.G[storage.TaskOwnerRecord](db).Where("task_id = ?", task.ID).Find(ctx)
	require.NoError(t, err)
	assert.Empty(t, owners)

	assignees, err = gorm.G[storage.TaskAssigneeRecord](db).Where("task_id = ?", task.ID).Find(ctx)
	require.NoError(t, err)
	assert.Empty(t, assignees)

	boards, err = gorm.G[storage.BoardTaskRecord](db).Where("task_id = ?", task.ID).Find(ctx)
	require.NoError(t, err)
	require.Len(t, boards, 1)
	assert.Nil(t, boards[0].ColumnID)
	assert.Equal(t, int64(5), boards[0].Position)

	tasks, err := gorm.G[storage.TaskRecord](db).Where("project_id = ?", project.ID).Find(ctx)
	require.NoError(t, err)
	assert.Len(t, tasks, 1, "update must not insert a second task")
}

func Test_TaskStorageRepository_Save_UnknownTask(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	repo := repositories.NewGormTaskStorageRepository(db)

	err := repo.Save(ctx, &models.Task{ID: 999999, Title: "ghost"})
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
}

func Test_TaskStorageRepository_AssignUser(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	repo := repositories.NewGormTaskStorageRepository(db)
	fixture := seedTask(t, ctx, db, repo)

	require.NoError(t, repo.AssignUser(ctx, fixture.task.ID, uint(fixture.user.ID)))
	require.NoError(t, repo.AssignUser(ctx, fixture.task.ID, uint(fixture.user.ID)), "assigning twice is a no-op")

	assignees, err := gorm.G[storage.TaskAssigneeRecord](db).Where("task_id = ?", fixture.task.ID).Find(ctx)
	require.NoError(t, err)
	require.Len(t, assignees, 1)
	assert.Equal(t, int64(fixture.user.ID), assignees[0].UserID)

	assert.Error(t, repo.AssignUser(ctx, 999999, uint(fixture.user.ID)), "unknown task is rejected by the foreign key")
}

func Test_TaskStorageRepository_Complete(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	repo := repositories.NewGormTaskStorageRepository(db)
	fixture := seedTask(t, ctx, db, repo)

	require.NoError(t, repo.Complete(ctx, fixture.task.ID))

	record, err := gorm.G[storage.TaskRecord](db).Where("id = ?", fixture.task.ID).First(ctx)
	require.NoError(t, err)
	assert.True(t, record.Completed)

	assert.ErrorIs(t, repo.Complete(ctx, 999999), gorm.ErrRecordNotFound)
}

func Test_TaskStorageRepository_MovePosition(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	repo := repositories.NewGormTaskStorageRepository(db)
	fixture := seedTask(t, ctx, db, repo)

	other := &models.Column{Title: "other column", BoardId: uint(fixture.board.ID)}
	require.NoError(t, db.WithContext(ctx).Create(other).Error)

	require.NoError(t, repo.MovePosition(ctx, fixture.task.ID, int64(fixture.board.ID), int64(other.ID), 0))

	placement, err := gorm.G[storage.BoardTaskRecord](db).Where("task_id = ?", fixture.task.ID).First(ctx)
	require.NoError(t, err)
	require.NotNil(t, placement.ColumnID)
	assert.Equal(t, int64(other.ID), placement.ColumnID)
	assert.Equal(t, int64(0), placement.Position, "position 0 is written, not skipped as a zero value")

	require.NoError(t, repo.MovePosition(ctx, fixture.task.ID, int64(fixture.board.ID), 0, 3))

	placement, err = gorm.G[storage.BoardTaskRecord](db).Where("task_id = ?", fixture.task.ID).First(ctx)
	require.NoError(t, err)
	assert.Nil(t, placement.ColumnID, "a zero column detaches the task from its column")
	assert.Equal(t, int64(3), placement.Position)

	err = repo.MovePosition(ctx, fixture.task.ID, 999999, int64(other.ID), 1)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound, "a board the task is not placed on")
}

type taskFixture struct {
	user  *models.Identity
	board *models.Board
	task  *models.Task
}

func seedTask(t *testing.T, ctx context.Context, db *gorm.DB, repo repositories.TaskStorageRepository) taskFixture {
	t.Helper()

	user := &models.Identity{FullName: "user", Email: "user@example.com"}
	require.NoError(t, db.WithContext(ctx).Create(user).Error)

	workspace := &models.Workspace{Title: "ws", Active: true}
	require.NoError(t, db.WithContext(ctx).Create(workspace).Error)
	project := &models.Project{Title: "project", WorkspaceID: workspace.ID}
	require.NoError(t, db.WithContext(ctx).Create(project).Error)
	board := &models.Board{Title: "board", ProjectID: project.ID}
	require.NoError(t, db.WithContext(ctx).Create(board).Error)
	column := &models.Column{Title: "column", BoardId: uint(board.ID)}
	require.NoError(t, db.WithContext(ctx).Create(column).Error)

	task := &models.Task{
		Title:     "task",
		ProjectId: int64(project.ID),
		Boards: []models.TaskBoard{
			{BoardId: int64(board.ID), ColumnId: int64(column.ID), Position: 1},
		},
	}
	require.NoError(t, repo.Save(ctx, task))

	return taskFixture{user: user, board: board, task: task}
}
