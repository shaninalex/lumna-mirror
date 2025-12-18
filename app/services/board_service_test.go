package services_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_ServiceBoard_ProjectBoards(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	boardB := tests.CreateBoard(ctx, projectA.GetId(), "B")
	boardC := tests.CreateBoard(ctx, projectA.GetId(), "C")

	projectB := tests.CreateProjectWithName(ctx, "ProjectB")
	_ = tests.CreateBoard(ctx, projectB.GetId(), "D")
	_ = tests.CreateBoard(ctx, projectB.GetId(), "E")
	_ = tests.CreateBoard(ctx, projectB.GetId(), "F")

	service := services.NewBoardService()

	boards, err := service.ProjectBoards(ctx, projectA.GetId())
	assert.NoError(t, err)
	assert.Equal(t, 3, len(boards))

	testBoardsNames := []string{boardA.Name, boardB.Name, boardC.Name}
	queriedBoards := make([]string, len(testBoardsNames))
	for i, u := range boards {
		queriedBoards[i] = u.Name
	}

	for _, e := range testBoardsNames {
		assert.Contains(t, queriedBoards, e, fmt.Sprintf("Board %s does not exists in queriedBoards %v", e, queriedBoards))
	}
}

func Test_ServiceBoard_Update(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	service := services.NewBoardService()
	newName := "B"

	err := service.Update(ctx, boardA.GetId(), db.Set(db.Field("name", newName)))
	assert.NoError(t, err)

	repo := repositories.NewBoardRepository()
	dbBoard, _ := repo.Get(ctx, boardA.GetId())
	assert.Equal(t, newName, dbBoard.Name)
}

func Test_ServiceBoard_Delete(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	service := services.NewBoardService()

	err := service.Delete(ctx, boardA.GetId())
	assert.NoError(t, err)

	repo := repositories.NewBoardRepository()
	dbBoard, err := repo.Get(ctx, boardA.GetId())
	assert.Nil(t, dbBoard)
	assert.Error(t, err)
}

func Test_ServiceBoard_Create(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	projectA := tests.CreateProjectWithName(ctx, "ProjectA")

	boardA := &models.Board{
		Name:      "A",
		ProjectId: projectA.GetId(),
	}

	service := services.NewBoardService()
	createdBoard, err := service.Create(ctx, boardA)
	assert.NoError(t, err)
	assert.Equal(t, boardA.Id, createdBoard.Id)
	assert.Equal(t, boardA.Name, createdBoard.Name)

	repo := repositories.NewBoardRepository()
	dbBoard, err := repo.Get(ctx, boardA.GetId())
	assert.NoError(t, err)
	assert.Equal(t, boardA.GetId(), dbBoard.GetId())
	assert.Equal(t, boardA.Name, dbBoard.Name)
}

func Test_ServiceBoard_Lists(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")

	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	listA := tests.CreateBoardList(ctx, boardA.GetId(), "listA")
	listB := tests.CreateBoardList(ctx, boardA.GetId(), "listB")
	listC := tests.CreateBoardList(ctx, boardA.GetId(), "listC")

	boardB := tests.CreateBoard(ctx, projectA.GetId(), "B")
	_ = tests.CreateBoardList(ctx, boardB.GetId(), "listD")
	_ = tests.CreateBoardList(ctx, boardB.GetId(), "listE")
	_ = tests.CreateBoardList(ctx, boardB.GetId(), "listF")

	service := services.NewBoardService()
	lists, err := service.Lists(ctx, boardA.GetId())
	assert.NoError(t, err)

	testListNames := []string{
		listA.Name,
		listB.Name,
		listC.Name,
	}
	queriedLists := make([]string, len(testListNames))
	for i, u := range lists {
		queriedLists[i] = u.Name
	}

	for _, e := range testListNames {
		assert.Contains(t, queriedLists, e, fmt.Sprintf("List %s does not exists in queriedLists %v", e, queriedLists))
	}
}

func Test_ServiceBoard_ListCreate(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	projectA := tests.CreateProjectWithName(ctx, "ProjectA")
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")

	service := services.NewBoardService()

	list := &models.BoardList{
		Name:    "A",
		BoardId: boardA.GetId(),
	}
	list, err := service.ListCreate(ctx, list)
	assert.NoError(t, err)
	assert.NotNil(t, list)
	assert.NotEqual(t, 0, list.GetId())

	repo := repositories.NewBoardListRepository()
	dbList, err := repo.Get(ctx, list.GetId())
	assert.NoError(t, err)
	assert.Equal(t, list.Name, dbList.Name)
}

func Test_ServiceBoard_ListUpdate(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	projectA := tests.CreateProject(ctx)
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	listA := tests.CreateBoardList(ctx, boardA.GetId(), "listA")

	service := services.NewBoardService()
	newName := "new name"

	err := service.ListUpdate(ctx, listA.GetId(), db.Set(db.Field("name", newName)))
	assert.NoError(t, err)

	repo := repositories.NewBoardListRepository()
	dbList, err := repo.Get(ctx, listA.GetId())
	assert.NoError(t, err)
	assert.Equal(t, newName, dbList.Name)
}

func Test_ServiceBoard_ListDelete(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	projectA := tests.CreateProject(ctx)
	boardA := tests.CreateBoard(ctx, projectA.GetId(), "A")
	listA := tests.CreateBoardList(ctx, boardA.GetId(), "listA")

	service := services.NewBoardService()

	err := service.ListDelete(ctx, listA.GetId())
	assert.NoError(t, err)

	repo := repositories.NewBoardListRepository()
	dbList, err := repo.Get(ctx, listA.GetId())
	assert.Error(t, err)
	assert.Nil(t, dbList)
}
