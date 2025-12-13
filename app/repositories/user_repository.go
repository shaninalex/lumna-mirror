package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/mail"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

var _ Repository[models.User] = (*UserRepository)(nil)

func (s *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	row := db.FromContext(ctx).QueryRow(`
		SELECT id, email, active, password_hash, created_at, updated_at FROM users WHERE email = ?
	`, email)
	err := row.Scan(&user.Id, &user.Email, &user.Active, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *UserRepository) GetUserPasswordHash(ctx context.Context, id uint) (string, error) {
	pwdHash := ""
	row := db.FromContext(ctx).QueryRow(`
		SELECT password_hash FROM users WHERE id = ?
	`, id)
	err := row.Scan(&pwdHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrorUserNotFound
		}
		return "", err
	}

	return pwdHash, nil
}

func (s *UserRepository) Get(ctx context.Context, id uint) (*models.User, error) {
	user := &models.User{}
	row := db.FromContext(ctx).QueryRow(`
		SELECT id, email, active, password_hash, created_at, updated_at FROM users WHERE id = ?
	`, id)
	err := row.Scan(&user.Id, &user.Email, &user.Active, &user.PasswordHash, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorUserNotFound
		}
		return nil, err
	}

	return user, nil
}

func (s *UserRepository) Delete(ctx context.Context, id uint) error {
	res, err := db.FromContext(ctx).ExecContext(ctx, `DELETE FROM users WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}

func (s *UserRepository) Create(ctx context.Context, user *models.User) error {
	query := `
		INSERT INTO users (email, active, created_at, updated_at, password_hash)
		VALUES (?, ?, ?, ?, ?)
	`
	user.UpdatedAt = time.Now()
	user.CreatedAt = time.Now()

	if err := validateUserObject(*user); err != nil {
		return err
	}

	result, err := db.FromContext(ctx).ExecContext(ctx, query, user.Email, user.Active, user.CreatedAt, user.UpdatedAt, user.PasswordHash)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.SetId(uint(id))
	return nil
}

func (s *UserRepository) List(ctx context.Context, opts ...db.Option) ([]*models.User, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT id, email, active, created_at, updated_at FROM users %s`,
		where,
	)

	rows, err := db.FromContext(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*models.User{}
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (s *UserRepository) Update(ctx context.Context, userID uint, opts ...db.Option) error {
	if len(opts) == 0 {
		return ErrorUserNoFieldsToUpdate
	}

	// always update updated_at
	opts = append(opts, db.Option{
		Key:   "updated_at",
		Value: time.Now(),
	})

	set, args := db.Set(opts)
	if set == "" {
		return ErrorUserNoFieldsToUpdate
	}

	query := fmt.Sprintf(
		`UPDATE users %s WHERE id = ?`,
		set,
	)

	args = append(args, userID)

	res, err := db.FromContext(ctx).ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrorNoRowsAffected
	}

	return nil
}

func (s *UserRepository) Count(ctx context.Context, opts ...db.Option) (int, error) {
	where, args := db.Where(opts)

	query := fmt.Sprintf(
		`SELECT count(*) FROM users %s`,
		where,
	)

	var count int
	row := db.FromContext(ctx).QueryRowContext(ctx, query, args...)
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func validateUserObject(user models.User) error {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return err
	}

	if user.PasswordHash == "" {
		return ErrorUserInvalidPassword
	}

	return nil
}
