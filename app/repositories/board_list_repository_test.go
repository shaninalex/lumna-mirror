package repositories_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_RepositoryBoardList_Create(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	listRepo := repositories.NewBoardListRepository()
	board := tests.CreateBoard(ctx, tests.CreateProject(ctx).GetId(), "test")
	entry := models.BoardList{
		Name:    "todo",
		BoardId: board.GetId(),
	}
	err := listRepo.Create(ctx, &entry)
	assert.NoError(t, err, "Should create board list without errors")
}

func Test_RepositoryBoardList_Count(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardListRepository()
	board := tests.CreateBoard(ctx, tests.CreateProject(ctx).GetId(), "test")
	_ = tests.CreateBoardList(ctx, board.Id, "test")

	count, err := repo.Count(ctx, nil)
	assert.NoError(t, err, "Should count board list without errors")
	assert.Equal(t, 1, count)

	count, err = repo.Count(ctx, db.Eq("name", "test"))
	assert.NoError(t, err, "Should board list without errors")
	assert.Equal(t, 1, count)

	count, err = repo.Count(ctx, db.Eq("name", "none existed board list"))
	assert.NoError(t, err, "Should board list without errors")
	assert.Equal(t, 0, count)
}

func Test_RepositoryBoardList_Delete(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardListRepository()
	board := tests.CreateBoard(ctx, tests.CreateProject(ctx).GetId(), "test")
	boardList := tests.CreateBoardList(ctx, board.Id, "test")

	err := repo.Delete(ctx, boardList.GetId())
	assert.NoError(t, err, "Should not find deleted board list")

	count, err := repo.Count(ctx, db.Eq("name", "test"))
	assert.NoError(t, err, "Should not find deleted board list")
	assert.Equal(t, 0, count)
}

func Test_RepositoryBoardList_Get(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardListRepository()
	board := tests.CreateBoard(ctx, tests.CreateProject(ctx).GetId(), "test")
	boardList := tests.CreateBoardList(ctx, board.Id, "test")
	dbBoardList, err := repo.Get(ctx, boardList.GetId())
	assert.NoError(t, err, "Should get board list by id without errors")
	assert.Equal(t, boardList.GetId(), dbBoardList.GetId())
	assert.Equal(t, boardList.Name, dbBoardList.Name)
	assert.Equal(t, boardList.Order, dbBoardList.Order)
}

func Test_RepositoryBoardList_List(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	board := tests.CreateBoard(ctx, tests.CreateProject(ctx).GetId(), "test")
	boardListA := tests.CreateBoardList(ctx, board.Id, "test A")
	boardListB := tests.CreateBoardList(ctx, board.Id, "test B")

	repo := repositories.NewBoardListRepository()
	boardLists, err := repo.List(ctx, nil)
	assert.NoError(t, err, "Should list projects without errors")
	assert.Equal(t, 2, len(boardLists))

	testBoardLists := []string{boardListA.Name, boardListB.Name}
	queriedBoardLists := make([]string, len(boardLists))
	for i, u := range boardLists {
		queriedBoardLists[i] = u.Name
	}

	for _, e := range testBoardLists {
		assert.Contains(t, queriedBoardLists, e, fmt.Sprintf("Board list %s does not exists in queriedBoardLists %v", e, queriedBoardLists))
	}
}

func Test_RepositoryBoardList_ListWithOption(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	board := tests.CreateBoard(ctx, tests.CreateProject(ctx).GetId(), "test")
	_ = tests.CreateBoardList(ctx, board.Id, "test A")
	_ = tests.CreateBoardList(ctx, board.Id, "test B")

	repo := repositories.NewBoardListRepository()
	boardLists, err := repo.List(ctx, db.Eq("name", "test A"))
	assert.NoError(t, err, "Should list projects without errors")
	assert.Equal(t, 1, len(boardLists))
	assert.Equal(t, "test A", boardLists[0].Name)
}

func Test_RepositoryBoardList_Update(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	board := tests.CreateBoard(ctx, tests.CreateProject(ctx).GetId(), "test")
	boardList := tests.CreateBoardList(ctx, board.Id, "test A")

	newName := "New Name"

	repo := repositories.NewBoardListRepository()
	err := repo.Update(ctx, boardList.GetId(), db.Set(db.Field("name", newName)))
	assert.NoError(t, err, "Should update project without errors")

	updatedProject, _ := repo.Get(ctx, boardList.GetId())
	assert.Equal(t, newName, updatedProject.Name)
}
