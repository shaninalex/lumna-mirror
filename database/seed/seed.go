package seed

import (
	"time"

	"github.com/google/uuid"
	"gitlab.com/shaninalex/jajirra/models"
	"gitlab.com/shaninalex/jajirra/models/builders"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB, userID uuid.UUID) error {
	now := time.Now()

	// Create organization
	org := models.Organization{
		ID:          uuid.New(),
		UserID:      userID,
		Title:       "Self Org",
		Description: "Organization for personal projects including Taskiro.",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := db.Create(&org).Error; err != nil {
		return err
	}

	// Attach user to this org
	if err := db.Model(&models.User{}).
		Where("id = ?", userID).
		Update("organization_id", org.ID).Error; err != nil {
		return err
	}

	// Create project "Taskiro"
	project := models.Project{
		ID:             uuid.New(),
		UserID:         userID,
		OrganizationID: org.ID,
		Title:          "Taskiro",
		ProjectKey:     "taskiro",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := db.Create(&project).Error; err != nil {
		return err
	}

	// Create project statuses ( board columns )
	statusTodo := builders.NewIssueStatusBuilder().
		Project(&project).ID(uuid.New()).ProjectID(project.GetID()).
		Title("Todo").Index(0).Build()
	statusInProgress := builders.NewIssueStatusBuilder().
		Project(&project).ID(uuid.New()).ProjectID(project.GetID()).
		Title("In Progress").Index(1).Build()
	statusTest := builders.NewIssueStatusBuilder().
		Project(&project).ID(uuid.New()).ProjectID(project.GetID()).
		Title("Testing").Index(2).Build()
	statusDone := builders.NewIssueStatusBuilder().
		Project(&project).ID(uuid.New()).ProjectID(project.GetID()).
		Title("Done").Index(3).Complete(true).Build()
	statuses := []*models.IssueStatus{statusTodo, statusInProgress, statusTest, statusDone}
	if err := db.Create(&statuses).Error; err != nil {
		return err
	}

	// Create sprints
	sprint1 := models.Sprint{
		ID:             uuid.New(),
		UserID:         userID,
		OrganizationID: org.ID,
		Title:          "Sprint 1: Foundation",
		Description:    "Set up the base of Taskiro: authentication, projects, issues",
		StartDate:      now,
		EndDate:        now.AddDate(0, 0, 14),
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	sprint2 := models.Sprint{
		ID:             uuid.New(),
		UserID:         userID,
		OrganizationID: org.ID,
		Title:          "Sprint 2: Task Management",
		Description:    "Implement issues, epics, and sprint features",
		StartDate:      now.AddDate(0, 0, 15),
		EndDate:        now.AddDate(0, 0, 30),
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	if err := db.Create(&[]models.Sprint{sprint1, sprint2}).Error; err != nil {
		return err
	}

	// Create epics
	epicUserMgmt := models.Epic{
		ID:          uuid.New(),
		UserID:      userID,
		ProjectID:   project.ID,
		Title:       "User Management",
		Description: "Epic covering authentication, user settings, and organization membership",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	epicCorePM := models.Epic{
		ID:          uuid.New(),
		UserID:      userID,
		ProjectID:   project.ID,
		Title:       "Core Project Management",
		Description: "Epic for projects, sprints, and issues",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	epicUI := models.Epic{
		ID:          uuid.New(),
		UserID:      userID,
		ProjectID:   project.ID,
		Title:       "UI/UX Improvements",
		Description: "Epic for frontend polish and usability",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	if err := db.Create(&[]models.Epic{epicUserMgmt, epicCorePM, epicUI}).Error; err != nil {
		return err
	}

	// Create issues
	issues := []models.Issue{
		{
			ID:             uuid.New(),
			UserID:         userID,
			EpicID:         &epicUserMgmt.ID,
			OrganizationID: org.ID,
			SprintID:       &sprint1.ID,
			ProjectID:      project.ID,
			Assignee:       "alex",
			Type:           "feature",
			Title:          "Use only Material SDK",
			Description:    "Replace default material components with SDK and manual created ui elements",
			IssueStatusID:  statusTodo.GetID(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID:             uuid.New(),
			UserID:         userID,
			EpicID:         &epicUserMgmt.ID,
			OrganizationID: org.ID,
			SprintID:       &sprint1.ID,
			ProjectID:      project.ID,
			Assignee:       "alex",
			Type:           "feature",
			Title:          "Implement authentication (login/register)",
			Description:    "Use Ory Kratos for identity management",
			IssueStatusID:  statusDone.GetID(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID:             uuid.New(),
			UserID:         userID,
			EpicID:         &epicUserMgmt.ID,
			OrganizationID: org.ID,
			SprintID:       &sprint1.ID,
			ProjectID:      project.ID,
			Assignee:       "alex",
			Type:           "feature",
			Title:          "Add user profile & settings page",
			Description:    "Allow users to update their information",
			IssueStatusID:  statusTodo.GetID(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID:             uuid.New(),
			UserID:         userID,
			EpicID:         &epicCorePM.ID,
			OrganizationID: org.ID,
			SprintID:       &sprint2.ID,
			ProjectID:      project.ID,
			Assignee:       "alex",
			Type:           "feature",
			Title:          "Create projects & organizations",
			Description:    "Implement CRUD for projects/organizations",
			IssueStatusID:  statusTest.GetID(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID:             uuid.New(),
			UserID:         userID,
			EpicID:         &epicCorePM.ID,
			OrganizationID: org.ID,
			SprintID:       &sprint2.ID,
			ProjectID:      project.ID,
			Assignee:       "alex",
			Type:           "feature",
			Title:          "Add issues & epics",
			Description:    "Core task tracking functionality",
			IssueStatusID:  statusInProgress.GetID(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
		{
			ID:             uuid.New(),
			UserID:         userID,
			EpicID:         &epicUI.ID,
			OrganizationID: org.ID,
			SprintID:       &sprint2.ID,
			ProjectID:      project.ID,
			Assignee:       "alex",
			Type:           "improvement",
			Title:          "Polish dashboard UI",
			Description:    "Make Taskiro visually appealing with Tailwind & animations",
			IssueStatusID:  statusTodo.GetID(),
			CreatedAt:      now,
			UpdatedAt:      now,
		},
	}
	if err := db.Create(&issues).Error; err != nil {
		return err
	}

	return nil
}
