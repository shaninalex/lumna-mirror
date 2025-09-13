// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_ProjectCreate(t *testing.T) {
	m := tdata.Manager()
	org, user, _ := tdata.CreatePack(m.Ctx)
	projectManager := domain.NewProjectManagement()
	title := uuid.NewString()
	projectDto := &dto.ProjectDto{
		Title: title,
	}
	result, err := projectManager.CreateProject(m.Ctx, user.GetID(), org.GetID(), projectDto)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.NotEqual(t, uuid.Nil, result.ID)
	assert.Equal(t, result.Title, title)
	assert.Equal(t, result.OrganizationID, org.GetID())
	assert.Equal(t, result.UserID, user.GetID())
}
