package org_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/jajirra/internal/apperrors"
	"gitlab.com/shaninalex/jajirra/internal/org"
	"gitlab.com/shaninalex/jajirra/tdata"
)

func Test_GetOrganizationByUserID(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)
	user := tdata.CreateUser(ctx)
	_org := tdata.CreateOrganisation(ctx, user)

	api := org.NewOrganizationApi()
	organization, err := api.Get(ctx, user.GetID())

	assert.NoError(t, err)
	assert.Equal(t, _org.GetID(), organization.GetID())
}

func Test_GetOrganizationNotFound(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)
	user := tdata.CreateUser(ctx)

	api := org.NewOrganizationApi()
	organization, err := api.Get(ctx, user.GetID())

	assert.Error(t, err)
	assert.ErrorIs(t, err, apperrors.OrgNotFound)
	assert.Nil(t, organization)
}
