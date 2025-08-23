package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/jajirra/database"
	"gitlab.com/shaninalex/jajirra/tdata"
)

func Test_UniqueUserCode(t *testing.T) {
	ctx := tdata.Ctx()
	tdata.Clear(ctx)
	tdata.CreateUser(ctx)

	code, err := database.GenerateUniqueUserCode(ctx, database.GetDB(ctx), 5)
	assert.NoError(t, err)

	code2, err := database.GenerateUniqueUserCode(ctx, database.GetDB(ctx), 5)
	assert.NoError(t, err)

	assert.NotEqual(t, code, code2)
}
