package tdata

import (
	"context"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/models"
)

func CreateOrganisation(ctx context.Context, user *models.User) *models.Organization {
	organization := models.NewOrganizationBuilder().
		User(*user).UserID(user.ID).
		Title(uuid.NewString()).
		Build()
	result := database.GetDB(ctx).Create(&organization)
	if result.Error != nil {
		panic(result.Error)
	}
	return organization
}

func CreateUser(ctx context.Context) *models.User {
	user := models.NewUserBuilder().ID(uuid.New()).Code(uuid.NewString()).Build()
	result := database.GetDB(ctx).Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

func CreateProject(ctx context.Context, org *models.Organization, user *models.User) *models.Project {
	project := models.NewProjectBuilder().
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
func CreatePack(ctx context.Context) (*models.Organization, *models.User, *models.Project) {
	user := CreateUser(ctx)
	org := CreateOrganisation(ctx, user)
	project := CreateProject(ctx, org, user)
	return org, user, project
}
