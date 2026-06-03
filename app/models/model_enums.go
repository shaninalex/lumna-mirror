package models

type TBodyFormat string

const (
	BodyFormatText     TBodyFormat = "TEXT"
	BodyFormatHTML     TBodyFormat = "HTML"
	BodyFormatMarkdown TBodyFormat = "MARKDOWN"
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

	// EmailStatusSkipped EmailStatusError dead end, executed with error
	EmailStatusSkipped EmailStatus = "skipped"
)
