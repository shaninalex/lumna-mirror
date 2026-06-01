package repositories

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type SprintRepository interface {
	List(ctx context.Context, query SprintListQuery) ([]*models.Sprint, error)
	GetByID(ctx context.Context, sprintID uint) (*models.Sprint, error)
	Create(ctx context.Context, sprint *models.Sprint) error
	Update(ctx context.Context, sprint *models.Sprint) error
	// Before implement this method - we should understand what exactly will be deleted
	// ScheduleDelete(ctx context.Context, SprintID uint) error
}

type GormSprintRepository struct {
	db *gorm.DB
}

func NewGormSprintRepository(db *gorm.DB) SprintRepository {
	return &GormSprintRepository{db: db}
}

type SprintListQuery struct {
	ProjectID *uint
}

func (s *GormSprintRepository) List(ctx context.Context, query SprintListQuery) ([]*models.Sprint, error) {
	var tasks []*models.Sprint
	db := s.db.WithContext(ctx)

	if query.ProjectID != nil {
		db = db.Where("project_id = ?", *query.ProjectID)
	}

	if err := db.Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil

}

func (s *GormSprintRepository) GetByID(ctx context.Context, SprintId uint) (*models.Sprint, error) {
	var Sprint models.Sprint
	if err := s.db.WithContext(ctx).
		Where("id = ?", SprintId).
		First(&Sprint).
		Error; err != nil {
		return nil, err
	}
	return &Sprint, nil
}

func (s *GormSprintRepository) Create(ctx context.Context, Sprint *models.Sprint) error {
	return s.db.WithContext(ctx).
		Create(Sprint).
		Error
}

func (s *GormSprintRepository) Update(ctx context.Context, sprint *models.Sprint) error {
	result := s.db.WithContext(ctx).Model(&models.Sprint{}).
		Where("id = ?", sprint.ID).
		Updates(sprint)
	return result.Error
}
