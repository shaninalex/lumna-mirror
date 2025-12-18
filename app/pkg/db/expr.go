package db

import "strings"

type Expr interface {
	build() (string, []any)
}

// Equal
type eq struct {
	field string
	value any
}

func Eq(field string, value any) Expr {
	return eq{field, value}
}

func (e eq) build() (string, []any) {
	return e.field + " = ?", []any{e.value}
}

// ILIKE
type ilike struct {
	field string
	value string
}

func ILike(field, value string) Expr {
	return ilike{field, value}
}

func (i ilike) build() (string, []any) {
	return i.field + " ILIKE ?", []any{i.value}
}

// IN
type in struct {
	field  string
	values []any
}

func In(field string, values []any) Expr {
	if len(values) == 0 {
		panic("db.In: empty values")
	}
	return in{field, values}
}

func (i in) build() (string, []any) {
	placeholders := make([]string, len(i.values))
	for idx := range placeholders {
		placeholders[idx] = "?"
	}

	return i.field + " IN (" + strings.Join(placeholders, ", ") + ")", i.values
}

// AND
type and struct {
	exprs []Expr
}

func And(exprs ...Expr) Expr {
	if len(exprs) == 0 {
		panic("db.And: no expressions")
	}
	return and{exprs}
}

func (a and) build() (string, []any) {
	var parts []string
	var args []any

	for _, e := range a.exprs {
		q, a2 := e.build()
		parts = append(parts, q)
		args = append(args, a2...)
	}

	return "(" + strings.Join(parts, " AND ") + ")", args
}

// OR
type or struct {
	exprs []Expr
}

func Or(exprs ...Expr) Expr {
	if len(exprs) == 0 {
		panic("db.Or: no expressions")
	}
	return or{exprs}
}

func (o or) build() (string, []any) {
	var parts []string
	var args []any

	for _, e := range o.exprs {
		q, a2 := e.build()
		parts = append(parts, q)
		args = append(args, a2...)
	}

	return "(" + strings.Join(parts, " OR ") + ")", args
}
