// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/utils"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/models/builders"
)

// CreateOrganisation - creates a new organisation.
func CreateOrganisation(ctx context.Context, user *models.User) *models.Organization {
	organization := builders.NewOrganizationBuilder().
		User(*user).UserID(user.ID).
		Title(uuid.NewString()).
		Build()
	result := database.GetDB(ctx).Create(&organization)
	if result.Error != nil {
		panic(result.Error)
	}
	return organization
}

// CreateUser - creates a new user.
func CreateUser(ctx context.Context) *models.User {
	user := builders.NewUserBuilder().
		ID(uuid.New()).
		Code(uuid.NewString()).
		Build()
	result := database.GetDB(ctx).Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	AddIdentity(&ory.Identity{
		Id: user.GetID().String(),
	})
	return user
}

// CreateProject - creates a new project.
func CreateProject(ctx context.Context, org *models.Organization, user *models.User) *models.Project {
	project := builders.NewProjectBuilder().
		User(*user).UserID(user.ID).
		Organization(*org).OrganizationID(org.ID).
		ProjectKey(utils.GenerateEntityCode("project")).
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

// CreateRandomStatuses - creates a new random statuses.
func CreateRandomStatuses(ctx context.Context, project *models.Project) []*models.TaskStatus {
	statuses := make([]*models.TaskStatus, 3)
	for i, stName := range []string{"Todo", "InProgress", "Done"} {
		st := models.TaskStatus{
			ProjectID: project.GetID(),
			Title:     stName,
			Complete:  false,
			Index:     0,
		}
		_db := database.GetDB(ctx)
		for _ = range 3 {
			task := builders.NewTaskBuilder().
				UserID(project.GetOwnerID()).
				OrganizationID(project.OrganizationID).
				ProjectID(project.GetID()).
				TaskStatusID(st.GetID()).
				TaskStatus(&st).
				Code(utils.GenerateEntityCode("task")).
				Title(uuid.NewString()).Build()
			st.Tasks = append(st.Tasks, task)
		}
		tx := _db.Save(&st)
		if tx.Error != nil {
			panic(tx.Error)
		}

		statuses[i] = &st
	}
	return statuses
}
