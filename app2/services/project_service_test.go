package services_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app2/models"
	"gitlab.com/shaninalex/lumna/app2/repositories"
	"gitlab.com/shaninalex/lumna/app2/services"
	"gitlab.com/shaninalex/lumna/app2/tests"
)

func Test_ServiceProject_CreateEmptyName(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	service := services.NewProjectManager()
	project := &models.Project{
		Name: "",
	}
	err := service.Create(ctx, project)
	assert.Error(t, err)
	assert.ErrorContains(t, err, "project name is required")
}

func Test_ServiceProject_Create(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	service := services.NewProjectManager()
	project := &models.Project{
		Name: "test",
	}
	err := service.Create(ctx, project)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, project.GetId())
	assert.Equal(t, "test", project.Name)

	repo := repositories.NewProjectRespository()
	dbProject, err := repo.Get(ctx, project.GetId())
	assert.NoError(t, err)
	assert.Equal(t, "test", dbProject.Name)
}

func Test_ServiceProject_List(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	service := services.NewProjectManager()
	projectA := &models.Project{Name: "test"}
	_ = service.Create(ctx, projectA)
	projectB := &models.Project{Name: "test 2"}
	_ = service.Create(ctx, projectB)

	projects, err := service.List(ctx)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(projects))

	testProjects := []string{projectA.Name, projectB.Name}
	queriedProjects := make([]string, len(projects))
	for i, u := range projects {
		queriedProjects[i] = u.Name
	}

	for _, e := range testProjects {
		assert.Contains(t, queriedProjects, e, fmt.Sprintf("Project %s does not exists in queriedProjects %v", e, queriedProjects))
	}
}
