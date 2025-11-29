package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/mail"
	"strings"
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

func (s *UserRepository) NewObject() models.User {
	return models.User{}
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
		return ErrorUserNoRowsAffected
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

func (s *UserRepository) List(ctx context.Context, options map[string]any) []*models.User {
	return []*models.User{}
}

func (s *UserRepository) Update(ctx context.Context, user *models.User, columns map[string]any) error {
	if len(columns) == 0 {
		return ErrorUserNoFieldsToUpdate
	}

	user.UpdatedAt = time.Now()
	columns["updated_at"] = user.UpdatedAt
	setParts := make([]string, 0, len(columns))
	args := make([]any, 0, len(columns)+1)

	for field, value := range columns {
		setParts = append(setParts, fmt.Sprintf("%s = ?", field))
		args = append(args, value)
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = ?",
		strings.Join(setParts, ", "),
	)
	args = append(args, user.Id)
	res, err := db.FromContext(ctx).ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println(err)
		return ErrorUserUnableToUpdate
	}

	if rowsAffected == 0 {
		return ErrorUserNoRowsAffected
	}

	// NOTE: that's not required second call, since we already know what fields will be updated
	updatedUser, err := s.Get(ctx, user.Id)
	*user = *updatedUser
	return err
}

func (s *UserRepository) Count(ctx context.Context, options map[string]any) (int, error) {
	return 0, nil
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
