package tdata

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
)

func CreateOrganisation(ctx context.Context, user *database.User) *database.Organization {
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

func CreateUser(ctx context.Context) *database.User {
	user := database.NewUserBuilder().ID(uuid.New()).Code(uuid.NewString()).Build()
	result := database.GetDB(ctx).Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

func CreateProject(ctx context.Context, org *database.Organization, user *database.User) *database.Project {
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

// CreatePack create a set of required fields
// - every user should be belonging to Organization
// - every project should be belonging to Organization
// - every organization should be created by some user
func CreatePack(ctx context.Context) (*database.Organization, *database.User, *database.Project) {
	user := CreateUser(ctx)
	org := CreateOrganisation(ctx, user)
	project := CreateProject(ctx, org, user)
	return org, user, project
}
