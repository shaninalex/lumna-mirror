package db

import (
	"context"
	"database/sql"

	"gitlab.com/shaninalex/lumna/app/global"
)

func FromContext(ctx context.Context) (conn *sql.DB) {
	conn = ctx.Value(global.ContextDB).(*sql.DB)
	if conn == nil {
		panic("database context is not set")
	}
	return conn
}
