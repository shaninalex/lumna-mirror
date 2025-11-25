package badge

import (
	"context"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

type BadgeReader interface {
	List(ctx context.Context, projectID int64) ([]*Badge, error)
}

type BadgeWriter interface {
	Create(ctx context.Context, badge *Badge) error
	Delete(ctx context.Context, projectID, badgeID int64) error
	AddToTask(ctx context.Context, taskID, badgeID int64) error
	DeleteFromTask(ctx context.Context, taskID, badgeID int64) error
}

type BadgeManager interface {
	BadgeReader
	BadgeWriter
}

type BadgeService struct {
}

func (b *BadgeService) AddToTask(ctx context.Context, taskID, badgeID int64) error {
	//TODO implement me
	panic("implement me")
}

func (b *BadgeService) DeleteFromTask(ctx context.Context, taskID, badgeID int64) error {
	//TODO implement me
	panic("implement me")
}

func NewBadgeService() *BadgeService {
	return &BadgeService{}
}

func (b *BadgeService) List(ctx context.Context, projectID int64) ([]*Badge, error) {
	badges, err := BadgeProjectList(ctx, db.GetDb(ctx), projectID)
	if err != nil {
		return nil, err
	}
	return badges, nil
}

func (b *BadgeService) Create(ctx context.Context, badge *Badge) error {
	if err := BadgeCreate(ctx, db.GetDb(ctx), badge); err != nil {
		return err
	}
	return nil
}

func (b *BadgeService) Delete(ctx context.Context, projectID, badgeID int64) error {
	return BadgeDelete(ctx, db.GetDb(ctx), projectID, badgeID)
}
