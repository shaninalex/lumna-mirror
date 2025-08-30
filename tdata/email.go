// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import "gitlab.com/shaninalex/flowreon/internal/email"

var _emailApi *TestEmailApi
var _emailStorage *TestEmailStorage

// GetTestEmailApi - returns the test email api.
func GetTestEmailApi() email.IEmailApi {
	if _emailApi == nil {
		_emailApi = &TestEmailApi{}
	}
	if _emailStorage == nil {
		_emailStorage = &TestEmailStorage{}
	}
	return _emailApi
}

// TestEmailApi - test email api.
type TestEmailApi struct{}

// SendVerificationEmail - send verification email.
func (s *TestEmailApi) SendVerificationEmail(token, to string) error {
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
