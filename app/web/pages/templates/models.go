package templates

type PageContext struct {
	Title string
	User  *PageUser
}

type PageUser struct {
	ID    string
	Email string
	Icon  string
}
