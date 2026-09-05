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
	ID           int
	Title        string
	Body         string
	Completed    bool
	Meta         string
	ProjectId    int
	OwnerId      int
	AssigneesIDs []int
	CreatedAt    time.Time
	UpdatedAt    time.Time

	Boards []TaskBoard
}

type TaskBoard struct {
	BoardId  int
	ColumnId int
	Position int
}

func (s *Task) String() string {
	return fmt.Sprintf("Task id=%d title=%s", s.ID, s.Title)
}

type TaskCreateOnBoard struct {
	Title     string `json:"title"`
	Body      string `json:"body"`
	Position  int    `json:"position"`
	BoardId   int    `json:"board_id"`
	ColumnId  int    `json:"column_id"`
	ProjectId int    `json:"project_id"`
}

type TaskCreateBacklog struct {
	Title string `json:"title"`
}

type RearangeTask struct {
	ColumnId int   `json:"column_id"`
	Tasks    []int `json:"tasks"`
}

type TransferTaskBetweenColumns struct {
	BoardId int          `json:"board_id"`
	From    RearangeTask `json:"from"`
	To      RearangeTask `json:"to"`
}
