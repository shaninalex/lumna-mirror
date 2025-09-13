// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package database_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/models/builders"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_UniqueUserCode(t *testing.T) {
	ctx := tdata.Manager().Ctx

	// Insert user with a known code
	existingCode := "testuser123"
	user := builders.NewUserBuilder().Code(existingCode).Build()
	result := database.GetDB(ctx).Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	// Try to generate a unique code
	code, err := database.GenerateUniqueUserCode(ctx, database.GetDB(ctx), 5)
	assert.NoError(t, err)
	assert.NotEqual(t, existingCode, code)
}
