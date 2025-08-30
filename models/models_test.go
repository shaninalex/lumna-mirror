// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package models_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/builders"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_ModelsBuilders(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)
	org, user, project := tdata.CreatePack(ctx)
	assert.NotNil(t, user)
	assert.NotNil(t, project)
	assert.Equal(t, org.UserID, user.ID)
	st := models.TaskStatus{
		ProjectID: project.GetID(),
		Title:     uuid.NewString(),
		Complete:  false,
		Index:     0,
	}
	database.GetDB(ctx).Save(&st)
	task := builders.NewTaskBuilder().
		User(*user).
		Organization(*org).
		Project(*project).
		Title(uuid.NewString()).
		Description(uuid.NewString()).
		TaskStatusID(st.GetID()).
		TaskStatus(&st).
		Build()
	result := database.GetDB(ctx).Create(&task)
	assert.NoError(t, result.Error)
}
