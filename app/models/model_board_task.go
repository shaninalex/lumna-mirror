package models

import (
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

const (
	EventTaskReorederd observer.Event = "TASK_REORDERED"
)

type BoardTask struct {
	BoardId  uint `gorm:"primaryKey" json:"board_id"`
	TaskId   uint `gorm:"not null" json:"task_id"`
	ColumnId uint `json:"column_id"`
	Position uint `json:"position"`
}
