package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_RepositoryTask_Get(t *testing.T) {
	ctx := tests.TestContext()
	repo := repositories.NewTaskRepository()

	pr := tests.CreateProject(ctx)
	br := tests.CreateBoard(ctx, pr.GetId(), "A")
	ls := tests.CreateBoardList(ctx, pr.GetId(), "A")

	task := &models.Task{
		BoardId: br.GetId(),
		ListId:  ls.GetId(),
		Name:    "A",
	}

	_ = repo.Create(ctx, task)

	dbTask, err := repo.Get(ctx, task.GetId())
	assert.Nil(t, err)
	assert.NotNil(t, dbTask)
	assert.Equal(t, "A", dbTask.Name)
}

func Test_RepositoryTask_Delete(t *testing.T) {
	ctx := tests.TestContext()
	pr := tests.CreateProject(ctx)
	br := tests.CreateBoard(ctx, pr.GetId(), "A")
	ls := tests.CreateBoardList(ctx, pr.GetId(), "A")
	task := tests.CreateTask(ctx, br.GetId(), ls.GetId(), "A")

	repo := repositories.NewTaskRepository()
	err := repo.Delete(ctx, task.GetId())
	assert.Nil(t, err)

	dbTask, err := repo.Get(ctx, task.GetId())
	assert.NotNil(t, err)
	assert.Nil(t, dbTask)
}

func Test_RepositoryTask_Create(t *testing.T) {
	ctx := tests.TestContext()
	repo := repositories.NewTaskRepository()

	pr := tests.CreateProject(ctx)
	br := tests.CreateBoard(ctx, pr.GetId(), "A")
	ls := tests.CreateBoardList(ctx, pr.GetId(), "A")

	task := &models.Task{
		BoardId: br.GetId(),
		ListId:  ls.GetId(),
		Name:    "A",
	}

	err := repo.Create(ctx, task)
	assert.NoError(t, err)
}

func Test_RepositoryTask_List(t *testing.T) {
	ctx := tests.TestContext()
	pr := tests.CreateProject(ctx)
	br := tests.CreateBoard(ctx, pr.GetId(), "A")
	ls := tests.CreateBoardList(ctx, pr.GetId(), "A")
	_ = tests.CreateTask(ctx, br.GetId(), ls.GetId(), "A")
	_ = tests.CreateTask(ctx, br.GetId(), ls.GetId(), "B")
	_ = tests.CreateTask(ctx, br.GetId(), ls.GetId(), "C")

	repo := repositories.NewTaskRepository()
	tasks, err := repo.List(ctx, db.Eq("board_id", br.GetId()))
	assert.Nil(t, err)
	assert.Len(t, tasks, 3)
}

func Test_RepositoryTask_Count(t *testing.T) {

}

func Test_RepositoryTask_Update(t *testing.T) {

}
