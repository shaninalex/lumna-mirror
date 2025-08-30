// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"context"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
)

var testKratosClient *TestKratosService

func NewTestKratosService() {
	testKratosClient = &TestKratosService{
		Identities: make([]*ory.Identity, 0),
	}
}

type TestKratosService struct {
	Identities []*ory.Identity
}

func (s *TestKratosService) Get(ctx context.Context, id uuid.UUID) (*ory.Identity, error) {
	for _, i := range s.Identities {
		if i.Id == id.String() {
			return i, nil
		}
	}
	return nil, apperrors.UserNotFound
}

func AddUser(identity *ory.Identity) {
	if testKratosClient == nil {
		NewTestKratosService()
	}
	testKratosClient.Identities = append(testKratosClient.Identities, identity)
}
