package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEq_Build(t *testing.T) {
	expr := Eq("active", true)

	sql, _ := expr.build()

	expected := "active = ?"
	assert.Equal(t, expected, sql)
}

func TestEq_Build_Table(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"simple", "active", "active = ?"},
		{"with_underscore", "created_at", "created_at = ?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr := Eq(tt.field, 1)
			sql, _ := expr.build()

			assert.Equal(t, tt.expected, sql)
		})
	}
}

func TestILike_Build(t *testing.T) {
	expr := ILike("email", "%@gmail.com%")

	sql, _ := expr.build()

	assert.Equal(t, "email ILIKE ?", sql)
}

func TestILike_Build_Table(t *testing.T) {
	tests := []struct {
		name     string
		field    string
		expected string
	}{
		{"simple", "email", "email ILIKE ?"},
		{"snake_case", "user_email", "user_email ILIKE ?"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr := ILike(tt.field, "%test%")
			sql, _ := expr.build()

			assert.Equal(t, tt.expected, sql)
		})
	}
}

func TestIn_Build(t *testing.T) {
	expr := In("id", []any{1, 2, 3})

	sql, _ := expr.build()

	assert.Equal(t, "id IN (?, ?, ?)", sql)
}

func TestIn_Build_Table(t *testing.T) {
	tests := []struct {
		name     string
		values   []any
		expected string
	}{
		{"single", []any{1}, "id IN (?)"},
		{"three", []any{1, 2, 3}, "id IN (?, ?, ?)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr := In("id", tt.values)
			sql, _ := expr.build()

			assert.Equal(t, tt.expected, sql)
		})
	}
}

func TestIn_EmptyValues_Panics(t *testing.T) {
	assert.Panics(t, func() {
		In("id", []any{})
	})
}

func TestAnd_Build(t *testing.T) {
	expr := And(
		Eq("active", true),
		ILike("email", "%@gmail.com%"),
	)

	sql, _ := expr.build()

	assert.Equal(t, "(active = ? AND email ILIKE ?)", sql)
}

func TestAnd_Build_Table(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{
			name: "two_conditions",
			expr: And(
				Eq("active", true),
				Eq("admin", false),
			),
			expected: "(active = ? AND admin = ?)",
		},
		{
			name: "three_conditions",
			expr: And(
				Eq("a", 1),
				Eq("b", 2),
				Eq("c", 3),
			),
			expected: "(a = ? AND b = ? AND c = ?)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, _ := tt.expr.build()
			assert.Equal(t, tt.expected, sql)
		})
	}
}

func TestAnd_NoExpressions_Panics(t *testing.T) {
	assert.Panics(t, func() {
		And()
	})
}

func TestOr_Build(t *testing.T) {
	expr := Or(
		Eq("active", true),
		ILike("email", "%@gmail.com%"),
	)

	sql, _ := expr.build()

	assert.Equal(t, "(active = ? OR email ILIKE ?)", sql)
}

func TestOr_Build_Table(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{
			name: "two_conditions",
			expr: Or(
				Eq("active", true),
				Eq("admin", false),
			),
			expected: "(active = ? OR admin = ?)",
		},
		{
			name: "three_conditions",
			expr: Or(
				Eq("a", 1),
				Eq("b", 2),
				Eq("c", 3),
			),
			expected: "(a = ? OR b = ? OR c = ?)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql, _ := tt.expr.build()
			assert.Equal(t, tt.expected, sql)
		})
	}
}

func TestOr_NoExpressions_Panics(t *testing.T) {
	assert.Panics(t, func() {
		Or()
	})
}

func TestComplex_AndOrILikeEq(t *testing.T) {
	expr := And(
		Eq("active", true),
		Or(
			ILike("email", "%@gmail.com%"),
			ILike("email", "%@yahoo.com%"),
		),
	)

	sql, _ := expr.build()

	assert.Equal(
		t,
		"(active = ? AND (email ILIKE ? OR email ILIKE ?))",
		sql,
	)
}

func TestComplex_AndInEq(t *testing.T) {
	expr := And(
		Eq("active", true),
		In("id", []any{1, 2, 3}),
	)

	sql, _ := expr.build()

	assert.Equal(
		t,
		"(active = ? AND id IN (?, ?, ?))",
		sql,
	)
}

func TestComplex_OrMixedOperators(t *testing.T) {
	expr := Or(
		Eq("admin", true),
		ILike("email", "%@company.com%"),
		In("id", []any{10, 20}),
	)

	sql, _ := expr.build()

	assert.Equal(
		t,
		"(admin = ? OR email ILIKE ? OR id IN (?, ?))",
		sql,
	)
}

func TestComplex_DeepNesting(t *testing.T) {
	expr := And(
		Eq("active", true),
		Or(
			And(
				ILike("email", "%@gmail.com%"),
				In("id", []any{1, 2}),
			),
			And(
				ILike("email", "%@yahoo.com%"),
				In("id", []any{3, 4}),
			),
		),
	)

	sql, _ := expr.build()

	assert.Equal(
		t,
		"(active = ? AND ((email ILIKE ? AND id IN (?, ?)) OR (email ILIKE ? AND id IN (?, ?))))",
		sql,
	)
}

func TestComplex_OrderIsPreserved(t *testing.T) {
	expr := And(
		Eq("a", 1),
		Eq("b", 2),
		Eq("c", 3),
	)

	sql, _ := expr.build()

	assert.Equal(
		t,
		"(a = ? AND b = ? AND c = ?)",
		sql,
	)
}
