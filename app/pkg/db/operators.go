package db

import (
	"fmt"
	"strings"
)

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

func Set(opts []Option) (string, []any) {
	if len(opts) == 0 {
		return "", nil
	}

	setParts := make([]string, 0, len(opts))
	args := make([]any, 0, len(opts))

	for _, opt := range opts {
		setParts = append(setParts, fmt.Sprintf("%s = ?", opt.Key))
		args = append(args, opt.Value)
	}

	return "SET " + strings.Join(setParts, ", "), args
}
