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

type TestManager struct {
	DB *gorm.DB
}

func newTestManager() {
	ctx = context.Background()
	db = InitTestDatabase()
	ctx = context.WithValue(ctx, base.ContextDB, db)
	m = &TestManager{}
}

func Manager() *TestManager {
	return m
}

func Ctx() context.Context {
	return ctx
}

func Config() base.IConfig {
	return config
}

func Clear(ctx context.Context) {
	clearDatabase(ctx)
}
