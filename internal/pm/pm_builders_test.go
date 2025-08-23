package pm_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/tdata"
)

func Test_ProjectList(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)

	org, user, project := tdata.CreatePack(ctx)

	assert.NotNil(t, user)
	assert.NotNil(t, project)
	assert.Equal(t, org.UserID, user.ID)
}

func Test_CreateIssue(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)
	org, user, project := tdata.CreatePack(ctx)
	issue := database.NewIssueBuilder().
		User(*user).
		Organization(*org).
		Project(*project).
		Title(uuid.NewString()).
		Description(uuid.NewString()).
		Status("todo").
		Type(database.IssueTypeTask).
		Build()
	result := database.GetDB(ctx).Create(&issue)
	assert.NoError(t, result.Error)
}
