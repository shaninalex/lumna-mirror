package services

import (
	"context"
	"encoding/json"
	"fmt"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg"
	"gitlab.com/shaninalex/lumna/app/pkg/bus"
	"gitlab.com/shaninalex/lumna/app/pkg/logger"
	"gorm.io/gorm"
)

type ProjectService struct {
	db       *gorm.DB
	eventBus bus.EventBus
	logger   logger.Logger
}

func NewProjectService(
	db *gorm.DB,
	eventBus bus.EventBus,
	logger logger.Logger,

) *ProjectService {
	s := &ProjectService{
		db:       db,
		eventBus: eventBus,
		logger:   logger,
	}
	s.init()
	return s
}

func (s *ProjectService) init() {
	s.eventBus.Subscribe(pkg.ProjectNewEvent, s.onNewProject)
}

func (s *ProjectService) onNewProject(ctx context.Context, data any) {
	p, ok := data.(*models.Project)
	if !ok {
		return
	}
	s.logger.Log(fmt.Sprintf("Project created: %s", p.Title))
	data, err := json.Marshal(p)
	if err != nil {
		s.logger.Log(fmt.Sprintf("Err: unable to parse project %s", err.Error()))
	}

	s.eventBus.Publish(ctx, pkg.EmailSendEvent, p)
}

func (s *ProjectService) Get(ctx context.Context, id uint) (*models.Project, error) {
	project := &models.Project{}
	if result := s.db.WithContext(ctx).Preload("Boards").Where("id = ?", id).First(&project); result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (s *ProjectService) List(ctx context.Context) ([]models.Project, error) {
	projects := []models.Project{}
	if result := s.db.WithContext(ctx).Preload("Boards").Find(&projects); result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (s *ProjectService) Create(ctx context.Context, title string, userID uint) (*models.Project, error) {
	project := &models.Project{
		Title: title,
	}
	if result := s.db.WithContext(ctx).Create(&project); result.Error != nil {
		return nil, result.Error
	}

	s.eventBus.Publish(ctx, pkg.ProjectNewEvent, project)
	return project, nil
}

func (s *ProjectService) Update(ctx context.Context, projectID uint, title string) (*models.Project, error) {
	database := s.db.WithContext(ctx)
	project := models.Project{}
	if result := database.Where("id = ?", projectID).First(&project); result.Error != nil {
		return nil, result.Error
	}

	project.Title = title

	if result := database.Save(&project); result.Error != nil {
		return nil, result.Error
	}

	return &project, nil
}

func (s *ProjectService) Delete(ctx context.Context, id uint) error {
	if result := s.db.WithContext(ctx).Where("id = ?", id).Delete(&models.Project{}); result.Error != nil {
		return result.Error
	}
	return nil
}
