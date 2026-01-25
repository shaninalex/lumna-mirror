package oauth

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	AuthCodeNotFound    = errors.New("authorization code not found")
	AuthCodeExpired     = errors.New("authorization code expired")
	AuthCodeAlreadyUsed = errors.New("authorization code already used")
)

type AuthorizationCode struct {
	Code                string
	ClientID            string
	UserID              string
	RedirectURI         string
	CodeChallenge       string
	CodeChallengeMethod string
	Scopes              []string
	ExpiresAt           time.Time
	Used                bool
}

type AuthorizationCodeStore interface {
	Save(ctx context.Context, code *AuthorizationCode) error
	Find(ctx context.Context, code string) (*AuthorizationCode, error)
	MarkUsed(ctx context.Context, code string) error
}

var _ AuthorizationCodeStore = (*InMemoryAuthorizationCodeStore)(nil)

type InMemoryAuthorizationCodeStore struct {
	mu    sync.Mutex
	codes map[string]*AuthorizationCode
}

func NewInMemoryAuthorizationCodeStore() *InMemoryAuthorizationCodeStore {
	return &InMemoryAuthorizationCodeStore{
		mu:    sync.Mutex{},
		codes: map[string]*AuthorizationCode{},
	}
}

func (s *InMemoryAuthorizationCodeStore) Save(ctx context.Context, code *AuthorizationCode) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.codes[code.Code] = code
	return nil
}

func (s *InMemoryAuthorizationCodeStore) Find(ctx context.Context, code string) (*AuthorizationCode, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	authCode, ok := s.codes[code]
	if !ok {
		return nil, AuthCodeNotFound
	}

	if time.Now().After(authCode.ExpiresAt) {
		return nil, AuthCodeExpired
	}

	if authCode.Used {
		return nil, AuthCodeAlreadyUsed
	}

	return authCode, nil
}

func (s *InMemoryAuthorizationCodeStore) MarkUsed(ctx context.Context, code string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	authCode, ok := s.codes[code]
	if !ok {
		return AuthCodeNotFound
	}

	if authCode.Used {
		return AuthCodeAlreadyUsed
	}

	authCode.Used = true
	return nil
}
