package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type PasswordCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (s *PasswordCredentials) Validate() error {
	validate := validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return nil
}

type EmailAuthProvider struct {
	db *gorm.DB
}

func NewLocalAuthProvider(db *gorm.DB) *EmailAuthProvider {
	s := &EmailAuthProvider{}
	return s
}

func (s *EmailAuthProvider) Name() string {
	return "local"
}

var UserNotFoundError = errors.New("user not found")

func (s *EmailAuthProvider) Authenticate(ctx context.Context, payload *PasswordCredentials) (*models.Identity, error) {
	credentials := models.Credential{}
	if result := s.db.WithContext(ctx).Preload("Identity").First(&credentials, "email = ?", payload.Email); result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, UserNotFoundError
		}
		return nil, result.Error
	}

	if credentials.PasswordHash == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := ValidatePassword(*credentials.PasswordHash, payload.Password); err != nil {
		return nil, err
	}

	return &credentials.Identity, nil
}
