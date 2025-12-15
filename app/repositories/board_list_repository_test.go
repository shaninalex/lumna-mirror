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

func Test_RepositoryBoardListCreate(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	listRepo := repositories.NewBoardListRepository()
	board := createBoard(ctx, createProject(ctx).GetId(), "test")
	entry := models.BoardList{
		Name:    "todo",
		BoardId: board.GetId(),
	}
	err := listRepo.Create(ctx, &entry)
	assert.NoError(t, err, "Should create board list without errors")
}

func Test_RepositoryBoardListCount(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardListRepository()
	board := createBoard(ctx, createProject(ctx).GetId(), "test")
	_ = createBoardList(ctx, board.Id, "test")

	count, err := repo.Count(ctx)
	assert.NoError(t, err, "Should count board list without errors")
	assert.Equal(t, 1, count)

	count, err = repo.Count(ctx, db.Option{Key: "name", Value: "test"})
	assert.NoError(t, err, "Should board list without errors")
	assert.Equal(t, 1, count)

	count, err = repo.Count(ctx, db.Option{Key: "name", Value: "none existed board list"})
	assert.NoError(t, err, "Should board list without errors")
	assert.Equal(t, 0, count)
}

func Test_RepositoryBoardListDelete(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardListRepository()
	board := createBoard(ctx, createProject(ctx).GetId(), "test")
	boardList := createBoardList(ctx, board.Id, "test")

	err := repo.Delete(ctx, boardList.GetId())
	assert.NoError(t, err, "Should not find deleted board list")

	count, err := repo.Count(ctx, db.Option{Key: "name", Value: "test"})
	assert.NoError(t, err, "Should not find deleted board list")
	assert.Equal(t, 0, count)
}

func Test_RepositoryBoardListGet(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardListRepository()
	board := createBoard(ctx, createProject(ctx).GetId(), "test")
	boardList := createBoardList(ctx, board.Id, "test")
	dbBoardList, err := repo.Get(ctx, boardList.GetId())
	assert.NoError(t, err, "Should get board list by id without errors")
	assert.Equal(t, boardList.GetId(), dbBoardList.GetId())
	assert.Equal(t, boardList.Name, dbBoardList.Name)
	assert.Equal(t, boardList.Order, dbBoardList.Order)
}

func Test_RepositoryBoardListList(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	board := createBoard(ctx, createProject(ctx).GetId(), "test")
	boardListA := createBoardList(ctx, board.Id, "test A")
	boardListB := createBoardList(ctx, board.Id, "test B")

	repo := repositories.NewBoardListRepository()
	boardLists, err := repo.List(ctx)
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

func Test_RepositoryBoardListListWithOption(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	board := createBoard(ctx, createProject(ctx).GetId(), "test")
	_ = createBoardList(ctx, board.Id, "test A")
	_ = createBoardList(ctx, board.Id, "test B")

	repo := repositories.NewBoardListRepository()
	boardLists, err := repo.List(ctx, db.Option{Key: "name", Value: "test A"})
	assert.NoError(t, err, "Should list projects without errors")
	assert.Equal(t, 1, len(boardLists))
	assert.Equal(t, "test A", boardLists[0].Name)
}

func Test_RepositoryBoardListUpdate(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	board := createBoard(ctx, createProject(ctx).GetId(), "test")
	boardList := createBoardList(ctx, board.Id, "test A")

	newName := "New Name"

	repo := repositories.NewBoardListRepository()
	err := repo.Update(ctx, boardList.GetId(), db.Option{Key: "name", Value: newName})
	assert.NoError(t, err, "Should update project without errors")

	updatedProject, _ := repo.Get(ctx, boardList.GetId())
	assert.Equal(t, newName, updatedProject.Name)
}
