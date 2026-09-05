package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type InvitationRepository interface {
	List(ctx context.Context) ([]models.Invitation, error)
	GetById(ctx context.Context, id int) (*models.Invitation, error)
	GetByHash(ctx context.Context, hash string) (*models.Invitation, error)
	Create(ctx context.Context, invitation *models.Invitation) error
	Update(ctx context.Context, invitation *models.Invitation) error
	Delete(ctx context.Context, id int) error
}

type GormInvitationRepository struct {
	db *gorm.DB
}

func NewGormInvitationRepository(db *gorm.DB) InvitationRepository {
	return &GormInvitationRepository{db: db}
}

func (r *GormInvitationRepository) List(ctx context.Context) ([]models.Invitation, error) {
	var invitations []models.Invitation

	err := r.db.WithContext(ctx).
		Find(&invitations).
		Error

	if err != nil {
		return nil, err
	}

	return invitations, nil
}

func (r *GormInvitationRepository) GetById(ctx context.Context, id int) (*models.Invitation, error) {
	var invitation models.Invitation

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&invitation).
		Error

	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func (r *GormInvitationRepository) GetByHash(ctx context.Context, hash string) (*models.Invitation, error) {
	var invitation models.Invitation

	err := r.db.WithContext(ctx).
		Where("token_hash = ?", hash).
		First(&invitation).
		Error

	if err != nil {
		return nil, err
	}

	return &invitation, nil
}

func (r *GormInvitationRepository) Create(ctx context.Context, invitation *models.Invitation) error {
	return r.db.WithContext(ctx).
		Create(invitation).
		Error
}

func (r *GormInvitationRepository) Update(ctx context.Context, invitation *models.Invitation) error {
	return r.db.WithContext(ctx).
		Save(invitation).
		Error
}

func (r *GormInvitationRepository) Delete(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Invitation{}).
		Error
}
