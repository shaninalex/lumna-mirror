// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package email

// IEmailApi - i email api.
type IEmailApi interface {
	SendVerificationEmail(token, to string) error
}
