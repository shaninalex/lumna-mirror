package adapter

import (
	"time"

	"gitlab.com/shaninalex/lumna/_old_app/domain"
)

// ChangeTaskStatusInput - change task status dto.
type ChangeTaskStatusInput struct {
	FromStatusID int64   `json:"from_status"`
	ToStatusID   int64   `json:"to_status"`
	FromIdx      int64   `json:"from_idx"`
	ToIdx        float64 `json:"to_idx"`
}

type TaskDetailInput struct {
	Title       string  `json:"title"`
	Completed   bool    `json:"completed"`
	Description string  `json:"description"`
	ListIndex   float64 `json:"list_index"`
	StatusID    int64   `json:"status_id"`
}

type TaskDto struct {
	Id          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	ProjectID   int64     `json:"project_id"`
	StatusID    int64     `json:"status_id"`
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	Description string    `json:"description"`
	ListIndex   float64   `json:"list_index"`
	Code        string    `json:"code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// related structures
	Badges   []*domain.Badge   `json:"badges"`
	Comments []*domain.Comment `json:"comments"`
}

func ToTaskDto(task *domain.Task) *TaskDto {
	out := &TaskDto{
		Id:        task.Id,
		UserID:    task.UserID,
		ProjectID: task.ProjectID,
		StatusID:  task.StatusID,
		Title:     task.Title,
		Completed: task.Completed,
		ListIndex: task.ListIndex,
		Code:      task.Code,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,

		Badges:   task.Badges,
		Comments: task.Comments,
	}

	if task.Description != nil {
		out.Description = *task.Description
	}

	return out
}
