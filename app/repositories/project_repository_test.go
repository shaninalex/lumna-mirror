package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/tests"
)

func Test_ProjectCreate(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewProjectRespository()

	project := models.Project{Name: "test"}
	err := repo.Create(ctx, &project)
	assert.NoError(t, err, "Should create project without errors")

	count, err := repo.Count(ctx)
	assert.NoError(t, err, "Should count project without errors")
	assert.Equal(t, 1, count)
}

func Test_ProjectCount(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewProjectRespository()

	_ = repo.Create(ctx, &models.Project{Name: "test"})
	_ = repo.Create(ctx, &models.Project{Name: "test 2"})

	count, err := repo.Count(ctx, db.Option{Key: "name", Value: "test 2"})
	assert.NoError(t, err, "Should count project without errors")
	assert.Equal(t, 1, count)

	count, err = repo.Count(ctx, db.Option{Key: "name", Value: "none existed project"})
	assert.NoError(t, err, "Should count project without errors")
	assert.Equal(t, 0, count)

	// We do not implement "OR" condition.
	// count, err = repo.Count(ctx, db.Option{Key: "name", Value: "test"}, db.Option{Key: "name", Value: "test 2"})
	// assert.NoError(t, err, "Should count project without errors")
	// assert.Equal(t, 2, count)
}

func Test_ProjectDelete(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewProjectRespository()
	projectA := models.Project{Name: "A"}
	_ = repo.Create(ctx, &projectA)
	projectB := models.Project{Name: "B"}
	_ = repo.Create(ctx, &projectB)

	count, _ := repo.Count(ctx)
	assert.Equal(t, 2, count)

	err := repo.Delete(ctx, projectA.GetId())
	assert.NoError(t, err, "Should delete project without errors")
	count, _ = repo.Count(ctx)
	assert.Equal(t, 1, count)

	count, _ = repo.Count(ctx, db.Option{Key: "name", Value: "A"})
	assert.Equal(t, 0, count)

	count, _ = repo.Count(ctx, db.Option{Key: "name", Value: "B"})
	assert.Equal(t, 1, count)
}

func Test_ProjectGet(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewProjectRespository()
	projectA := models.Project{Name: "A"}
	_ = repo.Create(ctx, &projectA)
	projectB := models.Project{Name: "B"}
	_ = repo.Create(ctx, &projectB)

	dbProjectA, err := repo.Get(ctx, projectA.GetId())
	assert.NoError(t, err, "Should get project without errors")
	assert.Equal(t, dbProjectA.Name, projectA.Name)

	dbProjectB, err := repo.Get(ctx, projectB.GetId())
	assert.NoError(t, err, "Should get project without errors")
	assert.Equal(t, dbProjectB.Name, projectB.Name)

	dbProjectC, err := repo.Get(ctx, 123)
	assert.Error(t, err, "Should NOT get non existed project")
	assert.Nil(t, dbProjectC)
}
