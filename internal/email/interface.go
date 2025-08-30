// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package email

type IEmailApi interface {
	SendVerificationEmail(token, to string) error
}
