package models

import (
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

const (
	EventTaskReorederd observer.Event = "TASK_REORDERED"
)

type BoardTask struct {
	BoardId  uint
	TaskId   uint
	ColumnId uint
	Position uint
}

func (s BoardTask) TableName() string {
	return "board_tasks"
}
