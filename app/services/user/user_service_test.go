package user_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/services/user"
	"gitlab.com/shaninalex/lumna/testutils"
)

func Test_IdentityService_Identity(t *testing.T) {
	ctx, db := testutils.SetupTest()
	defer testutils.ClearDB(db)

	_user := testutils.User(models.Identity{FullName: "test", Email: "test@test.com", Active: true}, db)
	service := user.NewUserService(repositories.NewGormIdentityRepository(db))
	idn, err := service.Identity(ctx, _user.ID)
	assert.NoError(t, err)
	assert.NotNil(t, idn)

	var identity models.Identity
	result := db.WithContext(ctx).Where("id = ?", idn.ID).First(&identity)
	assert.NoError(t, result.Error)
	assert.Equal(t, identity.ID, idn.ID)
}
