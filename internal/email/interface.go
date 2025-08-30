// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package email

// IEmailAPI - i email api.
type IEmailAPI interface {
	SendVerificationEmail(token, to string) error
}
