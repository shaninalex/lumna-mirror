package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type EmailRepository interface {
	ListPending(context.Context) []models.Email
	List(context.Context, map[string]any) []models.Email
	Get(context.Context, uint) (*models.Email, error)
	Update(context.Context, *models.Email) error
	Delete(context.Context, uint) error
	Create(context.Context, *models.Email) error
}

type GormEmailRepository struct {
	db *gorm.DB
}

func NewGormEmailRepository(db *gorm.DB) EmailRepository {
	return &GormEmailRepository{db: db}
}

// Create implements [EmailRepository].
func (g *GormEmailRepository) Create(ctx context.Context, email *models.Email) error {
	email.Status = models.EmailStatusPending
	return g.db.WithContext(ctx).Create(email).Error
}

// Delete implements [EmailRepository].
func (g *GormEmailRepository) Delete(ctx context.Context, id uint) error {
	return g.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Email{}).Error
}

// Get implements [EmailRepository].
func (g *GormEmailRepository) Get(ctx context.Context, id uint) (*models.Email, error) {
	var email models.Email
	if err := g.db.WithContext(ctx).Where("id = ?", id).First(&email).Error; err != nil {
		return nil, err
	}
	return &email, nil
}

// List implements [EmailRepository].
func (g *GormEmailRepository) List(ctx context.Context, filters map[string]any) []models.Email {
	var emails []models.Email
	query := g.db.WithContext(ctx)
	if len(filters) > 0 {
		query = query.Where(filters)
	}

	if err := query.Find(&emails).Error; err != nil {
		return []models.Email{}
	}

	return emails
}

// ListPending implements [EmailRepository].
func (g *GormEmailRepository) ListPending(ctx context.Context) []models.Email {
	var emails []models.Email
	if err := g.db.WithContext(ctx).Where("status = ?", "pending").Find(&emails).Error; err != nil {
		return []models.Email{}
	}
	return emails
}

// Update implements [EmailRepository].
func (g *GormEmailRepository) Update(ctx context.Context, email *models.Email) error {
	return g.db.WithContext(ctx).Model(&models.Email{}).Where("id = ?", email.ID).Save(email).Error
}
