package db

import (
	"fmt"
	"strings"
)

type Option struct {
	Key   string
	Value any
}

func Where(opts []Option) (string, []any) {
	if len(opts) == 0 {
		return "", nil
	}

	whereParts := make([]string, 0, len(opts))
	args := make([]any, 0, len(opts))

	for _, opt := range opts {
		whereParts = append(whereParts, fmt.Sprintf("%s = ?", opt.Key))
		args = append(args, opt.Value)
	}

	return "WHERE " + strings.Join(whereParts, " AND "), args
}
