package logger

import (
	"context"
	"fmt"
	"log"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ActivityService struct {
	db    *gorm.DB
	ctx   context.Context
	queue chan *models.ActivityLog
}

func ProvideActivityService(db *gorm.DB, ctx context.Context) *ActivityService {
	s := &ActivityService{
		db:    db,
		queue: make(chan *models.ActivityLog, 100),
		ctx:   ctx,
	}
	s.init()
	return s
}

func (s *ActivityService) init() {
	fmt.Println("[ActivityService] Start")
	go func() {
		for {
			select {
			case msg := <-s.queue:
				s.save(msg)
			case <-s.ctx.Done():
				fmt.Println("[Logger] Stopped")
				return
			}
		}
	}()
}

func (s *ActivityService) save(activity *models.ActivityLog) {
	if result := s.db.WithContext(s.ctx).Create(&activity); result.Error != nil {
		log.Println(result.Error)
	}
}

type ActivityLogPayload struct {
	Summary    string `json:"summary"`
	EntityID   uint   `json:"entity_id"`
	EntityType string `json:"entity_type"`
	Action     string `json:"action"`
}

func (s *ActivityService) Log(userId uint, activity *ActivityLogPayload) {
	s.queue <- &models.ActivityLog{
		Summary:    activity.Summary,
		EntityID:   activity.EntityID,
		EntityType: activity.EntityType,
		Action:     activity.Action,
		IdentityID: userId,
	}
}
