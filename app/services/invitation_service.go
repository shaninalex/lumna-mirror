package services

import (
	"context"
	"errors"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
	"gorm.io/gorm"
)

var (
	defaultValidUntil = 1 * time.Hour
)

type InvitationManager interface {
	Create(ctx context.Context, email, role string) (*models.Invitation, string, error)
	Accept(ctx context.Context, token string) error
	Delete(ctx context.Context, invitationId uint) error
	Reset(ctx context.Context, invitationId uint) (string, error)
	List(ctx context.Context) ([]models.Invitation, error)
}

func ProvideInvitationService(db *gorm.DB) *InvitationService {
	return &InvitationService{db: db}
}

type InvitationService struct {
	db *gorm.DB
}

func (s *InvitationService) List(ctx context.Context) ([]models.Invitation, error) {
	invitations := []models.Invitation{}
	if result := s.db.WithContext(ctx).Find(&invitations); result.Error != nil {
		return nil, result.Error
	}
	return invitations, nil
}

func (s *InvitationService) Create(ctx context.Context, email, role string) (*models.Invitation, string, error) {

	// token -> send in email, tokenHash -> save in db
	token, tokenHash, err := utils.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	invitation := &models.Invitation{
		Email:      email,
		Role:       role,
		State:      models.InvitationStatePending,
		TokenHash:  tokenHash,
		ValidUntil: time.Now().Add(defaultValidUntil),
	}
	if err := s.db.WithContext(ctx).Create(&invitation).Error; err != nil {
		return nil, "", err
	}
	return invitation, token, nil
}

func (s *InvitationService) Accept(ctx context.Context, token string) error {
	tokenHash := utils.HashToken(token)
	var invitation models.Invitation

	if err := s.db.WithContext(ctx).Where("token_hash = ?", tokenHash).
		First(&invitation).Error; err != nil {
		return err
	}

	if !invitation.IsValid() {
		return errors.New("invitation invalid or expired")
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// can create user
		invitation.State = models.InvitationStateAccepted
		if err := tx.Save(&invitation).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *InvitationService) Delete(ctx context.Context, invitationId uint) error {
	if result := s.db.WithContext(ctx).Where("invitation_id = ?", invitationId).Delete(&models.Invitation{}); result.RowsAffected == 0 {
		return errors.New("invitation not found")
	}
	return nil
}

func (s *InvitationService) Reset(ctx context.Context, invitationId uint) (string, error) {
	var invitation models.Invitation
	if err := s.db.WithContext(ctx).First(&invitation, invitationId).Error; err != nil {
		return "", err
	}

	token, tokenHash, err := utils.GenerateToken()
	if err != nil {
		return "", err
	}

	invitation.Reset(time.Now().Add(defaultValidUntil), tokenHash)
	if err := s.db.WithContext(ctx).Save(&invitation).Error; err != nil {
		return "", err
	}

	return token, nil
}
