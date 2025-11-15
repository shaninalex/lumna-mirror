// Model Badge.
//
// Badge can be applied to entity and change it behaviour. For example tasks with blocked badge can't be
// complete until block is resolved. Or tasks with "not now" can't be started. Or for example "waiting for approval",
// "fix needed". You can imagine that situations.
//
// Badge is different from the "tag". Tag - help organize lists, badge - define special behaviour. It's not a "status",
// it's a "state". Please, do not be confused.

package domain

import (
	"context"
	"encoding/json"
	"time"

	"gitlab.com/shaninalex/lumna/app/pkg/db"
)

type Badge struct {
	ID        int64       `json:"id"`
	ProjectID int64       `json:"project_id"`
	Title     string      `json:"title"`
	Config    BadgeConfig `json:"config"`
	CreatedAt time.Time   `json:"created_at"`
}

// BadgeConfig - badge config.
type BadgeConfig struct {
	Color string `json:"color,omitempty"`
}

// ToBadgeConfig - converts string to badge config.
func ToBadgeConfig(cnf string) BadgeConfig {
	var config BadgeConfig
	err := json.Unmarshal([]byte(cnf), &config)
	if err != nil {
		return NewBadgeStatusConfig()
	}
	return config
}

// NewBadgeStatusConfig - new task status config.
func NewBadgeStatusConfig() BadgeConfig {
	return BadgeConfig{
		Color: "default",
	}
}

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
