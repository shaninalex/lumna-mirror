package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetByID(ctx context.Context, id uint) (*models.Project, error)
	List(ctx context.Context) ([]models.Project, error)
	Create(ctx context.Context, project *models.Project) error
	Update(ctx context.Context, project *models.Project) error
	DeleteByID(ctx context.Context, id uint) error
}

type GormProjectRepository struct {
	db *gorm.DB
}

func NewGormProjectRepository(db *gorm.DB) ProjectRepository {
	return &GormProjectRepository{db: db}
}

func (r *GormProjectRepository) GetByID(ctx context.Context, id uint) (*models.Project, error) {
	var project models.Project

	err := r.db.WithContext(ctx).
		//Preload("Boards").
		Where("id = ?", id).
		First(&project).
		Error

	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *GormProjectRepository) List(ctx context.Context) ([]models.Project, error) {
	var projects []models.Project

	err := r.db.WithContext(ctx).
		//Preload("Boards").
		Find(&projects).
		Error

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *GormProjectRepository) Create(ctx context.Context, project *models.Project) error {
	return r.db.WithContext(ctx).
		Create(project).
		Error
}

func (r *GormProjectRepository) Update(ctx context.Context, project *models.Project) error {
	return r.db.WithContext(ctx).
		Model(&models.Project{}).
		Where("id = ?", project.ID).
		Save(project).
		Error
}

func (r *GormProjectRepository) DeleteByID(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.Project{}).
		Error
}
