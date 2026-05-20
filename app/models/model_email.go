package models

import (
	"time"
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
	ID         uint
	FromEmail  string
	ToEmail    string
	Subject    string
	Body       string
	Format     TBodyFormat
	SenderName string
	Headers    string
	Status     EmailStatus
	CC         string
	BCC        string
	ReplyTo    string
	SentAt     *time.Time
	CreatedAt  time.Time
}
