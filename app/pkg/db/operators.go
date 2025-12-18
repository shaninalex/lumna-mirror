package db

func Where(expr Expr) (string, []any) {
	if expr == nil {
		return "", nil
	}

	q, args := expr.build()
	return "WHERE " + q, args
}
