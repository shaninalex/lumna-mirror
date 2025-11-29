package repositories

import (
	"context"
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

func (s *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, bool) {
	return nil, false
}

func (s *UserRepository) NewObject() models.User {
	return models.User{}
}

func (s *UserRepository) Get(ctx context.Context, ID uint) (*models.User, bool) {
	return nil, false
}

func (s *UserRepository) Delete(ctx context.Context, userID uint) {

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
		return ErrorNoFieldsToUpdate
	}

	updates := ""
	for field, value := range columns {
		updates += fmt.Sprintf("%s = %v", field, value)
	}
	user.UpdatedAt = time.Now()
	updates += fmt.Sprintf("updated_at = %v", user.UpdatedAt)
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", updates)
	res, err := db.FromContext(ctx).ExecContext(ctx, query, user.Id)
	if err != nil {
		return err
	}

	c, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if c == 0 {
		return ErrorNoRowsAffected
	}

	return nil
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
		return ErrorInvalidPassword
	}

	return nil
}
