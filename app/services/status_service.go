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
	ListId uint   `json:"list_id"`
	Order  uint   `json:"order"`
	Title  string `json:"title"`
	// TODO: project id
}

func (s *statusService) Create(ctx context.Context, payload StatusUpdate) (*models.Status, error) {
	// NOTE: may be instead of ColumnUpdate use models.Column?
	list, err := s.listRepository.GetByID(ctx, payload.ListId)
	if err != nil {
		return nil, err
	}

	status := models.Status{
		Order:     payload.Order,
		Title:     payload.Title,
		ProjectID: list.ProjectID,
	}

	return s.statusRepository.Create(ctx, status)
}

func (s *statusService) Update(ctx context.Context, status *models.Status) (*models.Status, error) {
	return s.statusRepository.Update(ctx, status)
}

func (s *statusService) Delete(ctx context.Context, id uint) error {
	return s.statusRepository.Delete(ctx, id)
}
