package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/services/observer"
	"gitlab.com/shaninalex/lumna/testutils"
)

func Test_ProjectService_GetCode(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	bus := observer.ProvideEventBus()
	repo := repositories.NewGormProjectRepository(db)
	tRepo := repositories.NewGormTaskRepository(db)
	service := services.NewProjectService(repo, tRepo, bus)

	workspace := &models.Workspace{Title: "ws", Active: true}
	assert.NoError(t, db.WithContext(ctx).Create(workspace).Error)

	// project without meta -> first code is number 1
	noMeta := &models.Project{Title: "Project", Key: "PRO", WorkspaceID: workspace.ID}
	assert.NoError(t, repo.Create(ctx, noMeta))

	code, err := service.GetNewCode(ctx, noMeta.ID, "task")
	assert.NoError(t, err)
	assert.Equal(t, "PRO-1", code)
}
