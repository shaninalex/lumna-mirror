package db

import (
	"context"
	"database/sql"
)

const CONTEXT_DB = "database_context"

func FromContext(ctx context.Context) (conn *sql.DB) {
	conn = ctx.Value(CONTEXT_DB).(*sql.DB)
	if conn == nil {
		panic("database context is not set")
	}
	return conn
}
