// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/org/domain"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_GetOrganizationByUserID(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)
	user := tdata.CreateUser(ctx)
	_org := tdata.CreateOrganisation(ctx, user)

	api := domain.NewOrganizationApi()
	organization, err := api.Get(ctx, user.ID)

	assert.NoError(t, err)
	assert.Equal(t, _org.ID, organization.ID)
}

func Test_GetOrganizationNotFound(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)
	user := tdata.CreateUser(ctx)

	api := domain.NewOrganizationApi()
	organization, err := api.Get(ctx, user.ID)

	assert.Error(t, err)
	assert.ErrorIs(t, err, apperrors.OrgNotFound)
	assert.Nil(t, organization)
}
