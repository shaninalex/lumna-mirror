// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import "gitlab.com/shaninalex/flowreon/internal/email"

var _emailApi *TestEmailApi
var _emailStorage *TestEmailStorage

func GetTestEmailApi() email.IEmailApi {
	if _emailApi == nil {
		_emailApi = &TestEmailApi{}
	}
	if _emailStorage == nil {
		_emailStorage = &TestEmailStorage{}
	}
	return _emailApi
}

type TestEmailApi struct{}

func (s *TestEmailApi) SendVerificationEmail(token, to string) error {
	_emailStorage.Set(to, token)
	return nil
}

type TestEmailStorage struct {
	codes map[string]string
}

func (s *TestEmailStorage) Get(t string) string {
	if t, ok := s.codes[t]; ok {
		return t
	}
	return ""
}

func (s *TestEmailStorage) Set(key, v string) {
	s.codes[key] = v
}
