package models

import (
	"gitlab.com/shaninalex/lumna/app/services/observer"
)

const (
	EventTaskReorederd observer.Event = "TASK_REORDERED"
)

type BoardTask struct {
	BoardId  int
	TaskId   int
	ColumnId int
	Position int
}

func (s BoardTask) TableName() string {
	return "board_tasks"
}
