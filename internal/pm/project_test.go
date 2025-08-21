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
	tdata.Clear()
	ctx := tdata.Ctx()
	user := createUser(ctx)
	assert.NotNil(t, user)
}

func createUser(ctx context.Context) *database.User {
	user := database.User{
		ID:       uuid.New(),
		Settings: "",
	}
	result := database.GetDB(ctx).Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return &user
}
