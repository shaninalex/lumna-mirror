// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand/v2"
	"strings"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"gitlab.com/shaninalex/flowreon/internal/utils"
	"gitlab.com/shaninalex/flowreon/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (s *UserRepository) GetByField(ctx context.Context, db *sql.DB, field, value string) (*models.User, error) {
	user := &models.User{}
	query := fmt.Sprintf(`
	SELECT id, organization_id, email, settings, active, state, code, password_hash, created_at, updated_at
	FROM users WHERE %s = ? LIMIT 1
	`, field)
	row := db.QueryRowContext(ctx, query, value)
	err := row.Scan(&user.ID, &user.OrganizationID, &user.Email, &user.Settings, &user.Active, &user.State, &user.Code, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return user, nil
}

func (s *UserRepository) Save(ctx context.Context, db *sql.DB, user *models.User) (*models.User, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now()
	}
	user.ID = uuid.New()
	user.SetSettings(&models.DefaultUserSettings)
	user.Active = false
	user.State = models.UserStatePending
	user.Code = s.generateUniqueUserCode(ctx, db, 5)
	query := `
	INSERT INTO users (id, organization_id, email, settings, active, state, code, password_hash)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, &user.ID, &user.OrganizationID, &user.Email, &user.Settings, &user.Active, &user.State, &user.Code, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GenerateUniqueUserCode - generate unique user code.
func (s *UserRepository) generateUniqueUserCode(ctx context.Context, db *sql.DB, maxAttempts int) string {
	for i := 0; i < maxAttempts; i++ {
		b := strings.ToLower(randomdata.SillyName())
		code := fmt.Sprintf("%s%d", b, rand.IntN(100_000))

		_, err := s.GetByField(ctx, db, "code", code)
		if err != nil {
			return code
		}
		// Code exists. Retry
	}
	return utils.GenerateEntityCode("user")
}
