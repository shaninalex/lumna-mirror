package db

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/lumna/app/base"
)

func FromContext(ctx context.Context) (conn *sql.DB) {
	conn = ctx.Value(base.ContextDB).(*sql.DB)
	if conn == nil {
		panic("database context is not set")
	}
	return conn
}
