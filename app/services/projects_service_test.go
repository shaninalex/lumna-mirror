package services_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/testutils"
)

func Test_ProjectService_GetCode(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	repo := repositories.NewGormProjectRepository(db)

	workspace := &models.Workspace{Title: "ws", Active: true}
	assert.NoError(t, db.WithContext(ctx).Create(workspace).Error)

	// project without meta -> first code is number 1
	noMeta := &models.Project{Title: "Project", WorkspaceID: workspace.ID}
	assert.NoError(t, repo.Create(ctx, noMeta))

}
