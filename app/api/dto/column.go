package dto

import (
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
)

type ColumnDTO struct {
	ID        int64             `json:"id"`
	Title     string            `json:"title"`
	Meta      models.ColumnMeta `json:"meta"`
	BoardId   int64             `json:"board_id"`
	Position  int64             `json:"position"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt *time.Time        `json:"updated_at"`
}

func ColumnDTOToModel(c ColumnDTO) models.Column {
	return models.Column{
		ID:        c.ID,
		Title:     c.Title,
		Meta:      c.Meta,
		BoardId:   c.BoardId,
		Position:  c.Position,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}
