package models

import (
	"time"

	"gorm.io/datatypes"
)

type Email struct {
	ID         int               `json:"id"`
	FromEmail  string            `json:"from_email"`
	ToEmail    string            `json:"to_email"`
	Subject    string            `json:"subject"`
	Body       string            `json:"body"`
	Format     TBodyFormat       `json:"format"`
	SenderName string            `json:"sender_name"`
	Headers    string            `json:"headers"`
	Status     EmailStatus       `json:"status"`
	CC         string            `json:"cc"`
	BCC        string            `json:"bcc"`
	ReplyTo    string            `json:"reply_to"`
	SentAt     *time.Time        `json:"sent_at"`
	CreatedAt  time.Time         `json:"created_at"`
	Meta       datatypes.JSONMap `gorm:"type:json" json:"meta"`
}

func (s Email) GetMeta() map[string]any {
	return s.Meta
}
