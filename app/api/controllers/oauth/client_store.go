package oauth

import (
	"context"
	"errors"
	"slices"
)

type OAuthClient struct {
	ID            string // client_id
	Type          string // public | confidential
	RedirectURIs  []string
	AllowedScopes []string
}

// response_type=code
// client_id=angular-web
// redirect_uri=http://localhost:4200/auth/callback
// scope=openid profile
// state=xyz
// code_challenge=...
// code_challenge_method=S256

type ClientStore interface {
	FindByID(ctx context.Context, clientID string) (*OAuthClient, error)
}

var _ ClientStore = (*InMemoryClientStore)(nil)

func NewInMemoryClientStore() *InMemoryClientStore {
	s := &InMemoryClientStore{
		clients: make([]OAuthClient, 0),
	}
	s.init()
	return s
}

type InMemoryClientStore struct {
	clients []OAuthClient
}

func (s *InMemoryClientStore) init() {
	// Mock client for testing and devopment.
	// move this to config
	s.clients = append(s.clients, OAuthClient{
		ID:            "angular-web",
		Type:          "public",
		RedirectURIs:  []string{"http://localhost:4200/auth/callback"},
		AllowedScopes: []string{"openid", "profile"},
	})
}

var ClientNotFound = errors.New("client not found")

func (s *InMemoryClientStore) FindByID(ctx context.Context, clientID string) (*OAuthClient, error) {
	idx := slices.IndexFunc(s.clients, func(c OAuthClient) bool { return c.ID == clientID })
	if idx < 0 {
		return nil, ClientNotFound
	}
	return &s.clients[idx], nil
}
