package dto

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
)

type BoardTaskDto struct {
	BoardId  uint `json:"board_id"`
	ColumnId uint `json:"column_id"`
	Position uint `json:"position"`
}

func ToBoardTaskDto(bt models.BoardTask) BoardTaskDto {
	return BoardTaskDto{
		BoardId:  bt.BoardId,
		ColumnId: bt.ColumnId,
		Position: bt.Position,
	}
}

func ToBoardTaskDtoList(bts []models.BoardTask) []BoardTaskDto {
	result := []BoardTaskDto{}
	for _, bt := range bts {
		result = append(result, ToBoardTaskDto(bt))
	}
	return result
}

type TaskDto struct {
	ID        uint         `json:"id"`
	Title     string       `json:"title"`
	Body      string       `json:"body"`
	Completed bool         `json:"completed"`
	Meta      string       `json:"meta"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt *time.Time   `json:"updated_at"`
	Board     BoardTaskDto `json:"boards"`
}
