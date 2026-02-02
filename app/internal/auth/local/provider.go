package local

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"gitlab.com/shaninalex/lumna/app/internal/db"
	"gitlab.com/shaninalex/lumna/app/models"
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

type LocalAuthProvider struct {
}

func NewLocalAuthProvider() *LocalAuthProvider {
	s := &LocalAuthProvider{}
	return s
}

func (s *LocalAuthProvider) Name() string {
	return "local"
}

var UserNotFoundError = errors.New("user not found")

func (s *LocalAuthProvider) Authenticate(ctx context.Context, payload *PasswordCredentials) (*models.Identity, error) {
	database := db.GetDB(ctx)
	credentials := models.Credential{}
	if result := database.Preload("Identity").First(&credentials, "email = ?", payload.Email); result.Error != nil {
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
