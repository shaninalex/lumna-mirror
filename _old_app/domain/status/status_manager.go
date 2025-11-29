package status

import (
	"context"

	"gitlab.com/shaninalex/lumna/_old_app/pkg/db"
)

type StatusReader interface {
	Get(ctx context.Context, id int64) (*Status, error)
	ProjectStatuses(ctx context.Context, projectId int64) ([]*Status, error)
}

type StatusWriter interface {
	Create(ctx context.Context, projectId int64, title string) (*Status, error)
	Patch(ctx context.Context, data *Status) (*Status, error)
	Delete(ctx context.Context, statusId int64) error
	SortProjectStatus(ctx context.Context, data map[int64]int64) error
}

type StatusManager interface {
	StatusReader
	StatusWriter
}

type StatusService struct {
}

func (s StatusService) SortProjectStatus(ctx context.Context, data map[int64]int64) error {
	for idx, statusId := range data {
		status, err := TaskStatusByID(ctx, db.GetDb(ctx), statusId)
		if err != nil {
			return err
		}
		status.ListIndex = idx
		if err = TaskStatusUpdate(ctx, db.GetDb(ctx), status); err != nil {
			return err
		}
	}
	return nil
}

func (s StatusService) Get(ctx context.Context, id int64) (*Status, error) {
	status, err := TaskStatusByID(ctx, db.GetDb(ctx), id)
	if err != nil {
		return nil, err
	}
	return &Status{
		Id:        status.Id,
		Title:     status.Title,
		ListIndex: status.ListIndex,
		ProjectId: status.ProjectId,
		Config:    status.Config,
	}, nil
}

func (s StatusService) ProjectStatuses(ctx context.Context, projectId int64) ([]*Status, error) {
	dbStatuses, err := TaskStatusListByProject(ctx, db.GetDb(ctx), projectId)
	if err != nil {
		return nil, err
	}
	statuses := make([]*Status, len(dbStatuses))
	for i, status := range dbStatuses {
		statuses[i] = &Status{
			Id:        status.Id,
			Title:     status.Title,
			ListIndex: status.ListIndex,
			ProjectId: status.ProjectId,
			Config:    status.Config,
		}
	}
	return statuses, nil
}

func (s StatusService) Create(ctx context.Context, projectId int64, title string) (*Status, error) {
	status := &Status{
		ProjectId: projectId,
		Title:     title,
	}
	status, err := TaskStatusCreate(ctx, db.GetDb(ctx), status)
	if err != nil {
		return nil, err
	}
	return status, nil
}

func (s StatusService) Patch(ctx context.Context, data *Status) (*Status, error) {
	if err := TaskStatusUpdate(ctx, db.GetDb(ctx), data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s StatusService) Delete(ctx context.Context, statusId int64) error {
	return TaskStatusDelete(ctx, db.GetDb(ctx), statusId)
}

// NewStatusService - new status service
func NewStatusService() *StatusService {
	return &StatusService{}
}
