package services

import (
	"errors"
	"time"

	"gitlab.com/shaninalex/lumna/app/internal/utils"
	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

var (
	defaultValidUntil = 1 * time.Hour
)

type InvitationManager interface {
	Create(email, role string) (*models.Invitation, string, error)
	Accept(token string) error
	Delete(invitationId uint) error
	Reset(invitationId uint) (string, error)
}

type InvitationService struct {
	db *gorm.DB
}

func (s *InvitationService) Create(email, role string) (*models.Invitation, string, error) {

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
	if err := s.db.Create(&invitation).Error; err != nil {
		return nil, "", err
	}
	return invitation, token, nil
}

func (s *InvitationService) Accept(token string) error {
	tokenHash := utils.HashToken(token)
	var invitation models.Invitation

	if err := s.db.Where("token_hash = ?", tokenHash).
		First(&invitation).Error; err != nil {
		return err
	}

	if !invitation.IsValid() {
		return errors.New("invitation invalid or expired")
	}

	return s.db.Transaction(func(tx *gorm.DB) error {
		// can create user
		invitation.State = models.InvitationStateAccepted
		if err := tx.Save(&invitation).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *InvitationService) Delete(invitationId uint) error {
	if result := s.db.Where("invitation_id = ?", invitationId).Delete(&models.Invitation{}); result.RowsAffected == 0 {
		return errors.New("invitation not found")
	}
	return nil
}

func (s *InvitationService) Reset(invitationId uint) (string, error) {
	var invitation models.Invitation
	if err := s.db.First(&invitation, invitationId).Error; err != nil {
		return "", err
	}

	token, tokenHash, err := utils.GenerateToken()
	if err != nil {
		return "", err
	}

	invitation.Reset(time.Now().Add(defaultValidUntil), tokenHash)
	if err := s.db.Save(&invitation).Error; err != nil {
		return "", err
	}

	return token, nil
}
