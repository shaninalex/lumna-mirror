package domain

import (
	"github.com/google/uuid"
)

type ChangeTaskStatusDTO struct {
	FromStatusID uuid.UUID `json:"from_status"`
	ToStatusID   uuid.UUID `json:"to_status"`
	FromIdx      uint      `json:"from_idx"`
	ToIdx        uint      `json:"to_idx"`
}
