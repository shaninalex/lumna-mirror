package storage

import "time"

type TaskRecord struct {
	ID        int64     `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"not null"`
	Body      string    `gorm:"type:text"`
	Meta      string    `gorm:"type:text"`
	ProjectId int64     `gorm:"project_id"`
	Completed bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (s TaskRecord) TableName() string {
	return "tasks"
}

type TaskOwnerRecord struct {
	TaskID int64 `gorm:"primaryKey"`
	UserID int64 `gorm:"primaryKey"`
}

func (s TaskOwnerRecord) TableName() string {
	return "task_owners"
}

type TaskAssigneeRecord struct {
	TaskID int64 `gorm:"primaryKey"`
	UserID int64 `gorm:"primaryKey"`
}

func (s TaskAssigneeRecord) TableName() string {
	return "task_assignees"
}

type BoardTaskRecord struct {
	BoardID  int64 `gorm:"primaryKey"`
	TaskID   int64 `gorm:"primaryKey"`
	ColumnID int64 `gorm:"default:null"`
	Position int64 `gorm:"not null"`
}

func (s BoardTaskRecord) TableName() string {
	return "board_tasks"
}
