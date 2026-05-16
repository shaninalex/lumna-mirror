package services

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/repositories"
)

type StatusService struct {
	statusRepository repositories.StatusRepository
	listRepository  repositories.ListRepository
}

func NewStatusService(
	listRepository repositories.ListRepository,
	statusRepository repositories.StatusRepository,
) *StatusService {
	return &StatusService{
		listRepository:  listRepository,
		statusRepository: statusRepository,
	}
}

func (s *StatusService) Filter(ctx context.Context, listId uint) []models.Status {
	return s.statusRepository.FilterByList(ctx, listId)
}

func (s *StatusService) Get(ctx context.Context, statusID uint) (*models.Status, error) {
	return s.statusRepository.GetByID(ctx, statusID)
}

type StatusUpdate struct {
	ListId uint   `json:"list_id"`
	Order   uint   `json:"order"`
	Title   string `json:"title"`
}

func (s *StatusService) Create(ctx context.Context, payload StatusUpdate) (*models.Status, error) {
	// NOTE: may be instead of ColumnUpdate use models.Column?
	list, err := s.listRepository.GetByID(ctx, payload.ListId)
	if err != nil {
		return nil, err
	}

	status := models.Status{
		ListID:   list.ID,
		Order:     payload.Order,
		Title:     payload.Title,
		ProjectID: list.ProjectID,
	}

	return s.statusRepository.Create(ctx, status)
}

func (s *StatusService) Update(ctx context.Context, status *models.Status) (*models.Status, error) {
	return s.statusRepository.Update(ctx, status)
}

func (s *StatusService) Delete(ctx context.Context, id uint) error {
	return s.statusRepository.Delete(ctx, id)
}
