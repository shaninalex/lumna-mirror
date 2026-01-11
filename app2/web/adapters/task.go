package adapters

import (
	"time"

	"gitlab.com/shaninalex/lumna/app2/models"
)

type TaskDTO struct {
	Id        uint      `json:"id"`
	BoardId   uint      `json:"board_id"`
	ListId    uint      `json:"list_id"`
	Name      string    `json:"name"`
	Done      bool      `json:"done"`
	Order     uint      `json:"order"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToTaskDTO(u *models.Task) *TaskDTO {
	return &TaskDTO{
		Id:        u.Id,
		BoardId:   u.BoardId,
		ListId:    u.ListId,
		Name:      u.Name,
		Done:      u.Done,
		Order:     u.Order,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToTasksDTO(list []*models.Task) (tasks []*TaskDTO) {
	for _, task := range list {
		tasks = append(tasks, ToTaskDTO(task))
	}
	return tasks
}
