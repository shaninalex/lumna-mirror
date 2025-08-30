// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"gitlab.com/shaninalex/flowreon/internal/apperrors"
	"gitlab.com/shaninalex/flowreon/models"
)

// IdentityStorage - mock storage for ory identity
var IdentityStorage []*ory.Identity = make([]*ory.Identity, 0)

// SessionStorage - mock storage for ory session
var SessionStorage map[string]*ory.Session = map[string]*ory.Session{}

// AddIdentity - adds a new identity.
func AddIdentity(i *ory.Identity) {
	IdentityStorage = append(IdentityStorage, i)
}

// AddSession - adds a new session.
func AddSession(s models.AuthUser) string {
	var identity *ory.Identity
	for _, i := range IdentityStorage {
		if i.Id == s.GetID().String() {
			identity = i
			break
		}
	}

	if identity == nil {
		panic(apperrors.UserIdentityNotFound)
	}
	cookie := uuid.NewString()
	SessionStorage[cookie] = &ory.Session{
		Id:       uuid.NewString(),
		Identity: identity,
	}
	return cookie
}
