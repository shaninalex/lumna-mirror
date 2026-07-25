package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/services/logger"
	"gitlab.com/shaninalex/lumna/app/services/observer"
	"gitlab.com/shaninalex/lumna/testutils"
)

func Test_ProjectService_GetCode(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	bus := observer.ProvideEventBus()
	repo := repositories.NewGormProjectRepository(db)
	service := services.NewProjectService(repo, logger.ProvideLogger(ctx), bus)

	workspace := &models.Workspace{Title: "ws", Active: true}
	assert.NoError(t, db.WithContext(ctx).Create(workspace).Error)

	// project without meta -> first code is number 1
	noMeta := &models.Project{Title: "Project", Key: "PRO", WorkspaceID: workspace.ID}
	assert.NoError(t, repo.Create(ctx, noMeta))

	code, err := service.GetNewCode(ctx, noMeta.ID, "task")
	assert.NoError(t, err)
	assert.Equal(t, "PRO-1", code)

	// project whose meta holds the last used number yields the next one
	withMeta := &models.Project{Title: "Backend", Key: "BAC", WorkspaceID: workspace.ID}
	m := models.NewProjectMeta()
	m.LastEntityNumber["task"] = 5
	assert.NoError(t, withMeta.SetMeta(m))
	assert.NoError(t, repo.Create(ctx, withMeta))

	code, err = service.GetNewCode(ctx, withMeta.ID, "task")
	assert.NoError(t, err)
	assert.Equal(t, "BAC-6", code)

	// unknown entity type falls back to number 1
	code, err = service.GetNewCode(ctx, withMeta.ID, "epic")
	assert.NoError(t, err)
	assert.Equal(t, "BAC-1", code)

	// missing project surfaces an error
	_, err = service.GetNewCode(ctx, 999999, "task")
	assert.Error(t, err)
}
