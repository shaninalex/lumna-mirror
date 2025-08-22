package pm_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/tdata"
)

func Test_ProjectList(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)

	user := createUser(ctx)
	org := createOrganisation(ctx, user)
	project := createProject(ctx, org, user)

	assert.NotNil(t, user)
	assert.NotNil(t, project)
	assert.Equal(t, org.UserID, user.ID)
}

func createOrganisation(ctx context.Context, user *database.User) *database.Organization {
	organization := database.NewOrganizationBuilder().
		User(*user).UserID(user.ID).
		Title(uuid.NewString()).
		Build()
	result := database.GetDB(ctx).Create(&organization)
	if result.Error != nil {
		panic(result.Error)
	}
	return organization
}

func createUser(ctx context.Context) *database.User {
	user := database.NewUserBuilder().ID(uuid.New()).Build()
	result := database.GetDB(ctx).Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

func createProject(ctx context.Context, org *database.Organization, user *database.User) *database.Project {
	project := database.NewProjectBuilder().
		User(*user).UserID(user.ID).
		Organization(*org).OrganizationID(org.ID).
		Title(uuid.NewString()).
		Build()
	result := database.GetDB(ctx).Create(&project)
	if result.Error != nil {
		panic(result.Error)
	}
	return project
}
