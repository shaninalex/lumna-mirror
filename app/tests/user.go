package tests

import (
	"context"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

func CreateUser(ctx context.Context, email string) *models.User {
	h, _ := utils.CreatePasswordHash(email)
	user := models.User{
		Email:        email,
		PasswordHash: h,
	}

	query := `
		INSERT INTO users (email, active, created_at, updated_at, password_hash)
		VALUES (?, ?, ?, ?, ?)
	`
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()

	result, err := db.FromContext(ctx).ExecContext(ctx, query, user.Email, user.Active, user.CreatedAt, user.UpdatedAt, user.PasswordHash)
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	user.SetId(uint(id))
	return &user
}
