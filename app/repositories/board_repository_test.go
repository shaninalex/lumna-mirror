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

func Test_RepositoryBoard_Count(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardRepository()
	project := tests.CreateProject(ctx)
	entry := models.Board{
		Name:      "test",
		ProjectId: project.GetId(),
	}
	_ = repo.Create(ctx, &entry)

	count, err := repo.Count(ctx)
	assert.NoError(t, err, "Should count board without errors")
	assert.Equal(t, 1, count, "Should count 1 board in db")
}

func Test_RepositoryBoard_Create(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardRepository()
	project := tests.CreateProject(ctx)
	entry := models.Board{
		Name:      "test",
		ProjectId: project.GetId(),
	}
	err := repo.Create(ctx, &entry)
	assert.NoError(t, err, "Should create board without errors")
}

func Test_RepositoryBoard_Delete(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := tests.CreateProject(ctx)
	board := tests.CreateBoard(ctx, project.GetId(), "test")
	repo := repositories.NewBoardRepository()

	count, _ := repo.Count(ctx)
	assert.Equal(t, 1, count, "should have 1 board before deletion")

	err := repo.Delete(ctx, board.GetId())
	assert.NoError(t, err, "Should delete board without errors")

	count, _ = repo.Count(ctx)
	assert.Equal(t, 0, count, "should have 0 boards after deletion")
}

func Test_RepositoryBoard_Get(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := tests.CreateProject(ctx)
	board := tests.CreateBoard(ctx, project.GetId(), "test")
	repo := repositories.NewBoardRepository()

	dbBoard, err := repo.Get(ctx, board.GetId())
	assert.NoError(t, err, "Should get board without errors")
	assert.Equal(t, project.Name, dbBoard.Name)
	assert.Equal(t, project.GetId(), dbBoard.GetId())
}

func Test_RepositoryBoard_List(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := tests.CreateProject(ctx)
	boardA := tests.CreateBoard(ctx, project.GetId(), "A")
	boardB := tests.CreateBoard(ctx, project.GetId(), "B")
	repo := repositories.NewBoardRepository()

	boards, err := repo.List(ctx)
	assert.NoError(t, err, "Should list boards without errors")
	assert.Equal(t, 2, len(boards))

	testProjects := []string{boardA.Name, boardB.Name}
	queriedProjects := make([]string, len(boards))
	for i, u := range boards {
		queriedProjects[i] = u.Name
	}

	for _, e := range testProjects {
		assert.Contains(t, queriedProjects, e, fmt.Sprintf("Project %s does not exists in queriedProjects %v", e, queriedProjects))
	}
}

func Test_RepositoryBoard_ListWithOption(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := tests.CreateProject(ctx)
	boardA := tests.CreateBoard(ctx, project.GetId(), "A")
	_ = tests.CreateBoard(ctx, project.GetId(), "B")
	repo := repositories.NewBoardRepository()

	projects, err := repo.List(ctx, db.Option{Key: "name", Value: boardA.Name})
	assert.NoError(t, err, "Should list projects without errors")
	assert.Equal(t, 1, len(projects))
	assert.Equal(t, boardA.Name, projects[0].Name)
}

func Test_RepositoryBoard_Update(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := tests.CreateProject(ctx)
	boardA := tests.CreateBoard(ctx, project.GetId(), "A")
	repo := repositories.NewBoardRepository()

	newName := "New Name"

	err := repo.Update(ctx, boardA.GetId(), db.Option{Key: "name", Value: newName})
	assert.NoError(t, err, "Should update project without errors")

	updatedProject, _ := repo.Get(ctx, boardA.GetId())
	assert.Equal(t, newName, updatedProject.Name)
}
