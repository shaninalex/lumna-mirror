package email

// IEmailAPI - i email api.
type IEmailAPI interface {
	SendVerificationEmail(token, to string) error
}
