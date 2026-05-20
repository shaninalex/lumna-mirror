package models

type TBodyFormat string

const (
	BodyFormatText     TBodyFormat = "TEXT"
	BodyFormatHTML     TBodyFormat = "HTML"
	BodyFormatMarkdown TBodyFormat = "MARKDOWN"
)
