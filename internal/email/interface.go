package email

type IEmailApi interface {
	SendVerificationEmail(token, to string) error
}
