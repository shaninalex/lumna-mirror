package internal

type Event string

var (
	ProjectNewEvent Event = "project/new"
	EmailSendEvent  Event = "email/send"
)
