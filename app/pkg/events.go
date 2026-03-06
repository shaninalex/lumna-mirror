package pkg

type Event string

var (
	ProjectNewEvent Event = "project/new"
	EmailSendEvent  Event = "email/send"
)
