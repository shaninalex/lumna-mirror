package models

import (
	"fmt"
	"time"

	"gitlab.com/shaninalex/lumna/app/services/observer"
)

const (
	EventTaskCreated   observer.Event = "TASK_CREATED"
	EventTaskUpdated   observer.Event = "TASK_UPDATED"
	EventTaskCompleted observer.Event = "TASK_COMPLETED"
)

type Task struct {
	ID        int64
	Title     string
	Body      string
	Completed bool
	Meta      string

	ProjectId int64

	Boards       []TaskBoard
	OwnerId      int64
	AssigneesIDs []int64

	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskBoard struct {
	BoardId  int64
	ColumnId int64
	Position int64
}

func (s *Task) String() string {
	return fmt.Sprintf("Task id=%d title=%s", s.ID, s.Title)
}

type TaskCreateOnBoard struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	Position *uint  `json:"position"`
	BoardId  *uint  `json:"board_id"`
	ColumnId *uint  `json:"column_id"`
}

type TaskCreateBacklog struct {
	Title string `json:"title"`
}
