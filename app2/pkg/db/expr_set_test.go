package db

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetExpr_Build_SingleField(t *testing.T) {
	set := Set(Field("name", "Alice"))

	sql, args := set.Build()

	assert.Equal(t, "SET name = ?", sql)
	assert.Equal(t, []any{"Alice"}, args)
}

func TestSetExpr_Build_MultipleFields(t *testing.T) {
	set := Set(
		Field("name", "Alice"),
		Field("active", true),
	)

	sql, args := set.Build()

	assert.Equal(t, "SET name = ?, active = ?", sql)
	assert.Equal(t, []any{"Alice", true}, args)
}

func TestSetExpr_AppendField(t *testing.T) {
	set := Set(Field("name", "Alice"))
	set.Append(Field("updated_at", time.Now()))

	sql, args := set.Build()

	// SQL string should include both fields
	assert.Contains(t, sql, "name = ?")
	assert.Contains(t, sql, "updated_at = ?")
	// args should have length 2
	assert.Len(t, args, 2)
}

func TestSetExpr_OrderPreserved(t *testing.T) {
	set := Set(
		Field("a", 1),
		Field("b", 2),
	)
	set.Append(Field("c", 3))

	sql, args := set.Build()

	assert.Equal(t, "SET a = ?, b = ?, c = ?", sql)
	assert.Equal(t, []any{1, 2, 3}, args)
}

func TestSetExpr_EmptySet(t *testing.T) {
	// creating an empty SetExpr is allowed; Build() should return "SET " with no fields
	set := Set()
	sql, args := set.Build()

	assert.Equal(t, "SET ", sql)
	assert.Empty(t, args)
}
