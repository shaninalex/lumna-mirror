package utils

type BasePage struct {
	Title     string
	User      *PageUser
	CsrfToken string
}

type PageUser struct {
	ID    string
	Email string
	Icon  string
}
