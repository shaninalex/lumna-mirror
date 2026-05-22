package models

import (
	"time"

	"gorm.io/datatypes"
)

type EmailStatus string

var (
	// EmailStatusPending - waiting for execution
	EmailStatusPending EmailStatus = "pending"

	// EmailStatusRunning - currently executing
	EmailStatusRunning EmailStatus = "running"

	// EmailStatusSuccess - completed without errors
	EmailStatusSuccess EmailStatus = "success"

	// EmailStatusRepeat after N attempt Email becomes with error start
	EmailStatusRepeat EmailStatus = "repeat"

	// EmailStatusError dead end, executed with error
	EmailStatusError EmailStatus = "error"

	// EmailStatusError dead end, executed with error
	EmailStatusSkipped EmailStatus = "skipped"
)

type Email struct {
	ID         uint              `json:"id"`
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
