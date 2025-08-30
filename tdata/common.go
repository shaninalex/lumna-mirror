// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"
	"sync"

	"gitlab.com/shaninalex/flowreon/internal/base"
	"gorm.io/gorm"
)

var ctx context.Context
var m *TestManager
var db *gorm.DB
var lock = sync.RWMutex{}
var config = newTestConfig()

func init() {
	newTestManager()
}

// TestManager - test manager.
type TestManager struct {
	DB *gorm.DB
}

func newTestManager() {
	ctx = context.Background()
	db = InitTestDatabase()
	ctx = context.WithValue(ctx, base.ContextDB, db)
	m = &TestManager{
		DB: db,
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

// Config - config.
func Config() base.IConfig {
	return config
}

// Clear - clear.
func Clear(ctx context.Context) {
	clearDatabase(ctx)
}
