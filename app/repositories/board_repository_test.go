package repositories_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/tests"
)

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

func Test_RepositoryBoardCount(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardRepository()
	project := createProject(ctx)
	entry := models.Board{
		Name:      "test",
		ProjectId: project.GetId(),
	}
	_ = repo.Create(ctx, &entry)

	count, err := repo.Count(ctx)
	assert.NoError(t, err, "Should count board without errors")
	assert.Equal(t, 1, count, "Should count 1 board in db")
}

func Test_RepositoryBoardCreate(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewBoardRepository()
	project := createProject(ctx)
	entry := models.Board{
		Name:      "test",
		ProjectId: project.GetId(),
	}
	err := repo.Create(ctx, &entry)
	assert.NoError(t, err, "Should create board without errors")
}

func Test_RepositoryBoardDelete(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := createProject(ctx)
	board := createBoard(ctx, project.GetId(), "test")
	repo := repositories.NewBoardRepository()

	count, _ := repo.Count(ctx)
	assert.Equal(t, 1, count, "should have 1 board before deletion")

	err := repo.Delete(ctx, board.GetId())
	assert.NoError(t, err, "Should delete board without errors")

	count, _ = repo.Count(ctx)
	assert.Equal(t, 0, count, "should have 0 boards after deletion")
}

func Test_RepositoryBoardGet(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := createProject(ctx)
	board := createBoard(ctx, project.GetId(), "test")
	repo := repositories.NewBoardRepository()

	dbBoard, err := repo.Get(ctx, board.GetId())
	assert.NoError(t, err, "Should get board without errors")
	assert.Equal(t, project.Name, dbBoard.Name)
	assert.Equal(t, project.GetId(), dbBoard.GetId())
}

func Test_RepositoryBoardList(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := createProject(ctx)
	boardA := createBoard(ctx, project.GetId(), "A")
	boardB := createBoard(ctx, project.GetId(), "B")
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

func Test_RepositoryBoardListWithOption(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := createProject(ctx)
	boardA := createBoard(ctx, project.GetId(), "A")
	_ = createBoard(ctx, project.GetId(), "B")
	repo := repositories.NewBoardRepository()

	projects, err := repo.List(ctx, db.Option{Key: "name", Value: boardA.Name})
	assert.NoError(t, err, "Should list projects without errors")
	assert.Equal(t, 1, len(projects))
	assert.Equal(t, boardA.Name, projects[0].Name)
}

func Test_RepositoryBoardUpdate(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	project := createProject(ctx)
	boardA := createBoard(ctx, project.GetId(), "A")
	repo := repositories.NewBoardRepository()

	newName := "New Name"

	err := repo.Update(ctx, boardA.GetId(), db.Option{Key: "name", Value: newName})
	assert.NoError(t, err, "Should update project without errors")

	updatedProject, _ := repo.Get(ctx, boardA.GetId())
	assert.Equal(t, newName, updatedProject.Name)
}
