package storage

import "time"

type TaskRecord struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"not null"`
	Body      string    `gorm:"type:text"`
	Meta      string    `gorm:"type:text"`
	ProjectId int       `gorm:"project_id"`
	Completed bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (s TaskRecord) TableName() string {
	return "tasks"
}

type TaskOwnerRecord struct {
	TaskID int `gorm:"primaryKey"`
	UserID int `gorm:"primaryKey"`
}

func (s TaskOwnerRecord) TableName() string {
	return "task_owners"
}

type TaskAssigneeRecord struct {
	TaskID int `gorm:"primaryKey"`
	UserID int `gorm:"primaryKey"`
}

func (s TaskAssigneeRecord) TableName() string {
	return "task_assignees"
}

type BoardTaskRecord struct {
	BoardID  int `gorm:"primaryKey"`
	TaskID   int `gorm:"primaryKey"`
	ColumnID int `gorm:"default:null"`
	Position int `gorm:"not null"`
}

func (s BoardTaskRecord) TableName() string {
	return "board_tasks"
}
