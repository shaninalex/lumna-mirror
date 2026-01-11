package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
)

func TestWhere_NilExpr(t *testing.T) {
	sql, args := db.Where(nil)
	assert.Equal(t, "", sql)
	assert.Nil(t, args)
}

func TestWhere_Build(t *testing.T) {
	expr := db.And(
		db.Eq("active", true),
		db.ILike("email", "%@gmail.com%"),
	)

	sql, args := db.Where(expr)

	assert.Equal(t, "WHERE (active = ? AND email ILIKE ?)", sql)
	assert.Equal(t, []any{true, "%@gmail.com%"}, args)
}

func TestWhere_Table(t *testing.T) {
	tests := []struct {
		name    string
		expr    db.Expr
		expSQL  string
		expArgs []any
	}{
		{
			name:    "nil_expr",
			expr:    nil,
			expSQL:  "",
			expArgs: nil,
		},
		{
			name: "simple_and",
			expr: db.And(
				db.Eq("active", true),
				db.ILike("email", "%@gmail.com%"),
			),
			expSQL:  "WHERE (active = ? AND email ILIKE ?)",
			expArgs: []any{true, "%@gmail.com%"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, args := db.Where(tt.expr)
			assert.Equal(t, tt.expSQL, sql)
			assert.Equal(t, tt.expArgs, args)
		})
	}
}
