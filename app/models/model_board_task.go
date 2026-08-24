package models

import (
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

const (
	EventTaskReorederd observer.Event = "TASK_REORDERED"
)

type BoardTask struct {
	BoardId  uint   `gorm:"primaryKey" json:"board_id"`
	TaskId   string `gorm:"not null" json:"task_id"`
	ColumnId string `json:"column_id"`
	Position bool   `json:"position"`
}
