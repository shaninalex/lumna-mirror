// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import "gitlab.com/shaninalex/flowreon/internal/email"

var _emailAPI *TestEmailAPI
var _emailStorage *TestEmailStorage

// GetTestEmailAPI - returns the test email api.
func GetTestEmailAPI() email.IEmailAPI {
	if _emailAPI == nil {
		_emailAPI = &TestEmailAPI{}
	}
	if _emailStorage == nil {
		_emailStorage = &TestEmailStorage{}
	}
	return _emailAPI
}

// TestEmailAPI - test email api.
type TestEmailAPI struct{}

// SendVerificationEmail - send verification email.
func (s *TestEmailAPI) SendVerificationEmail(token, to string) error {
	_emailStorage.Set(to, token)
	return nil
}

// TestEmailStorage - test email storage.
type TestEmailStorage struct {
	codes map[string]string
}

// Get - returns the value.
func (s *TestEmailStorage) Get(t string) string {
	if t, ok := s.codes[t]; ok {
		return t
	}
	return ""
}

// Set - sets the value.
func (s *TestEmailStorage) Set(key, v string) {
	s.codes[key] = v
}
