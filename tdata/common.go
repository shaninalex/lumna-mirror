// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"
	"os"
	"sync"

	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/base"
	"gitlab.com/shaninalex/flowreon/internal/kratos"
	"gorm.io/gorm"
)

var ctx context.Context
var m *TestManager
var db *gorm.DB
var lock = sync.RWMutex{}

func init() {
	newTestManager()
}

// TestManager - test manager.
type TestManager struct {
	Ctx    context.Context
	DB     *gorm.DB
	Kratos kratos.IKratos
}

func newTestManager() {
	ctx = context.Background()
	conf := base.NewConfig(os.Getenv("CONFIG_PATH"))
	url := database.BuildDSN(conf)
	db = database.InitDB(url)
	ctx = context.WithValue(ctx, base.ContextDB, db)
	m = &TestManager{
		DB:     db,
		Ctx:    ctx,
		Kratos: NewMockKratosService(),
	}
}

// Manager - manager.
func Manager() *TestManager {
	return m
}

// Ctx - ctx.
func Ctx() context.Context {
	return ctx
}

// Clear - clear.
func Clear(ctx context.Context) {
	clearDatabase(ctx)
}
