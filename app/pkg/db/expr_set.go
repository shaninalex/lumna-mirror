package db

import (
	"strings"
)

type SetExpr struct {
	exprs []setField
}

type setField struct {
	field string
	value any
}

func Field(name string, value any) setField {
	return setField{
		field: name,
		value: value,
	}
}

func Set(fields ...setField) SetExpr {
	return SetExpr{
		exprs: fields,
	}
}

func (s *SetExpr) Append(field setField) {
	s.exprs = append(s.exprs, field)
}

func (s *SetExpr) Build() (string, []any) {
	var parts []string
	var args []any

	for _, f := range s.exprs {
		parts = append(parts, f.field+" = ?")
		args = append(args, f.value)
	}

	return "SET " + strings.Join(parts, ", "), args
}
