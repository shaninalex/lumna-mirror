// Copyright © 2025 Lumna. All rights reserved.

package email

// IEmailAPI - i email api.
type IEmailAPI interface {
	SendVerificationEmail(token, to string) error
}
