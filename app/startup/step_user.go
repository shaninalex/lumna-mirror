// Copyright © 2025 Lumna. All rights reserved.

package startup

import (
	"context"
	"database/sql"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/internal/db"
	"golang.org/x/crypto/bcrypt"
)

type StepUser struct {
	db    *sql.DB
	title string
}

func NewStepUser(db *sql.DB) *StepUser {
	return &StepUser{
		db:    db,
		title: "User",
	}
}

func (s *StepUser) SetTitle(title string) {
	s.title = title
}

func (s *StepUser) Execute() error {
	fmt.Printf("Run step: %s\n", s.title)

	fmt.Print("Enter your email: ")
	var email string
	fmt.Scanln(&email)

	fmt.Print("Enter password: ")
	var password string
	fmt.Scanln(&password)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := &db.User{
		Email:        email,
		PasswordHash: string(hash),
	}
	ctx := context.Background()
	user, err = db.UserSave(ctx, s.db, user)
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully registered", email)
	fmt.Println("Application started. You can now login.")

	return nil
}
