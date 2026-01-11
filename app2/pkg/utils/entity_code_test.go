package utils_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app2/pkg/utils"
)

func Test_EntityCode(t *testing.T) {
	code := utils.GenerateEntityCode("entity")

	parts := strings.Split(code, "-")
	assert.Equal(t, "entity", parts[0])

	numbersPartA, err := strconv.Atoi(parts[1])
	assert.NoError(t, err)

	code = utils.GenerateEntityCode("entity")
	parts = strings.Split(code, "-")
	numbersPartB, err := strconv.Atoi(parts[1])

	assert.NotEqual(t, numbersPartA, numbersPartB)
}
