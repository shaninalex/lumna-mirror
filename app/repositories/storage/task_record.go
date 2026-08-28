package storage

import "time"

type TaskRecord struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"not null"`
	Body      string    `gorm:"type:text"`
	Meta      string    `gorm:"type:text"`
	Completed bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type TaskOwnerRecord struct {
	TaskID int64 `gorm:"primaryKey"`
	UserID int64 `gorm:"primaryKey"`
}

type TaskAssigneeRecord struct {
	TaskID int64 `gorm:"primaryKey"`
	UserID int64 `gorm:"primaryKey"`
}

type BoardTaskRecord struct {
	BoardID  int64  `gorm:"primaryKey"`
	TaskID   int64  `gorm:"primaryKey"`
	ColumnID *int64 `gorm:"default:null"`
	Position int64  `gorm:"not null"`
}
