// Copyright © 2025 Lumna. All rights reserved.

package startup

import (
	"database/sql"
)

type Step interface {
	SetTitle(title string)
	Execute() error
}

type AppInitializer interface {
	Run() error
	RegisterStep(step Step)
}

type Initializer struct {
	steps []Step
}

func NewStartup(db *sql.DB) *Initializer {
	s := &Initializer{
		steps: []Step{},
	}
	s.RegisterStep(NewStepUser(db))
	return s
}

func (s *Initializer) RegisterStep(step Step) {
	s.steps = append(s.steps, step)
}

func (s *Initializer) Run() error {
	for _, step := range s.steps {
		if err := step.Execute(); err != nil {
			return err
		}
	}
	return nil
}
