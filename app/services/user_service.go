package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserManager interface {
	GetUser(ctx context.Context, userID uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	CheckPassword(ctx context.Context, userId uint, rawPassword string) error
	CreateUser(ctx context.Context, email string, rawPassword string) (*models.User, error)
	Update(ctx context.Context, userId uint, option ...repositories.Option) error
}

var _ UserManager = &UserService{}

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService() *UserService {
	return &UserService{}
}

// GetUser get user
func (s UserService) GetUser(ctx context.Context, userId uint) (*models.User, error) {
	user, err := s.repository.Get(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail get user by email
func (s UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CheckPassword check password
func (s UserService) CheckPassword(ctx context.Context, userId uint, password string) error {
	pwdHash, err := s.repository.GetUserPasswordHash(ctx, userId)
	if err != nil {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(password)); err != nil {
		return err
	}
	return nil
}

// CreateUser create user
func (s UserService) CreateUser(ctx context.Context, email, rawPassword string) (*models.User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: string(hash),
	}

	err = s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &models.User{
		Id:        user.Id,
		Email:     user.Email,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Update update user
func (s UserService) Update(ctx context.Context, userId uint, options ...repositories.Option) error {
	user, err := s.repository.Get(ctx, userId)
	if err != nil {
		return err
	}
	if err = s.repository.Update(ctx, user, options...); err != nil {
		return err
	}
	return nil
}
