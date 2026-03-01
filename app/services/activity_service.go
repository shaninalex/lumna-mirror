package services

import (
	"context"
	"log"

	"gitlab.com/shaninalex/lumna/app/models"
	"gorm.io/gorm"
)

type ActivityService struct {
	db *gorm.DB
}

func ProvideActivityService(db *gorm.DB) *ActivityService {
	return &ActivityService{
		db: db,
	}
}

func (s *ActivityService) Log(ctx context.Context, activity *models.ActivityLog) {
	if result := s.db.WithContext(ctx).Create(&activity); result.Error != nil {
		log.Println(result.Error)
	}
}
