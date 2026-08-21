package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type StatusService interface {
	Filter(ctx context.Context, listId uint) []models.Status
	Get(ctx context.Context, statusID uint) (*models.Status, error)
	Create(ctx context.Context, payload StatusUpdate) (*models.Status, error)
	Update(ctx context.Context, status *models.Status) (*models.Status, error)
	Delete(ctx context.Context, id uint) error
}

type statusService struct {
	statusRepository repositories.StatusRepository
	listRepository   repositories.ListRepository
}

func NewStatusService(
	listRepository repositories.ListRepository,
	statusRepository repositories.StatusRepository,
) StatusService {
	return &statusService{
		listRepository:   listRepository,
		statusRepository: statusRepository,
	}
}

func (s *statusService) Filter(ctx context.Context, listId uint) []models.Status {
	return s.statusRepository.FilterByList(ctx, listId)
}

func (s *statusService) Get(ctx context.Context, statusID uint) (*models.Status, error) {
	return s.statusRepository.GetByID(ctx, statusID)
}

type StatusUpdate struct {
	Title     string `json:"title"`
	Order     uint   `json:"order"`
	ListId    uint   `json:"list_id"`
	ProjectId uint   `json:"project_id"`
}

func (s *statusService) Create(ctx context.Context, payload StatusUpdate) (*models.Status, error) {
	status := models.Status{
		Order:     payload.Order,
		Title:     payload.Title,
		ListId:    payload.ListId,
		ProjectID: payload.ProjectId,
	}
	return s.statusRepository.Create(ctx, status)
}

func (s *statusService) Update(ctx context.Context, status *models.Status) (*models.Status, error) {
	return s.statusRepository.Update(ctx, status)
}

func (s *statusService) Delete(ctx context.Context, id uint) error {
	return s.statusRepository.Delete(ctx, id)
}
