package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/_old_app/pkg/utils"
)

// UserGetByField get user by field
func UserGetByField(ctx context.Context, db *sql.DB, field string, value any) (*User, error) {
	user := &User{}
	query := fmt.Sprintf(`
	SELECT id, email, settings, active, state, code, password_hash, created_at, updated_at
	FROM users WHERE %s = ? LIMIT 1
	`, field)
	row := db.QueryRowContext(ctx, query, value)
	err := row.Scan(&user.Id, &user.Email, &user.Settings, &user.Active, &user.State, &user.Code, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}

	return user, nil
}

// UserSave save user
func UserSave(ctx context.Context, db *sql.DB, user *User) (*User, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now()
	}
	user.SetSettings(&DefaultUserSettings)
	user.Active = false
	user.State = "pending"
	user.Code = utils.GenerateEntityCode("user")
	query := `
	INSERT INTO users (email, settings, code, password_hash)
	VALUES (?, ?, ?, ?)
	`
	_, err := db.ExecContext(ctx, query, &user.Email, &user.Settings, &user.Code, &user.PasswordHash)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserUpdate update user
func UserUpdate(ctx context.Context, db *sql.DB, user *User) error {
	query := `
		UPDATE users
		SET 
		    email = ?, 
		    settings = ?, 
		    active = ?, 
		    state = ?, 
		    code = ?,
			updated_at = ?
		WHERE id = ?
	`
	_, err := db.ExecContext(ctx, query,
		user.Email,
		user.Settings,
		user.Active,
		user.State,
		user.Code,
		user.UpdatedAt,
	)
	return err
}
