package testutils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/testutils"
)

func Test_TestDBConnection(t *testing.T) {
	db := testutils.ProvideTestDB()
	err := testutils.Migrate(db)
	assert.NoError(t, err)
}
