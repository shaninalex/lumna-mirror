// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"
	"database/sql"
	"os"
	"sync"

	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/database"
)

var lock = sync.RWMutex{}

// TestManager - test manager.
type TestManager struct {
	Ctx context.Context
	DB  *sql.DB
}

// Manager - manager.
func Manager() *TestManager {
	ctx := context.Background()
	conf := base.NewConfig(os.Getenv("CONFIG_PATH"))
	url := database.BuildDSN(conf)
	db, _ := database.InitDB(url).DB()
	ctx = context.WithValue(ctx, base.ContextDB, db)

	lock.Lock()
	resetDatabase(ctx)
	lock.Unlock()

	return &TestManager{
		DB:  db,
		Ctx: ctx,
	}
}
