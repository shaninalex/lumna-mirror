// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler_test

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/google/uuid"
	ory "github.com/ory/kratos-client-go"
	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/auth/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/auth/domain"
	"gitlab.com/shaninalex/flowreon/apps/auth/dto"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_AuthHooks_HandleHookRegister(t *testing.T) {
	m := tdata.Manager()
	r := web.DefaultRouter(m.DB, "test_router")
	identity := &ory.Identity{
		Id: uuid.NewString(),
		Traits: &models.UserTraits{
			Email: "test@test.com",
			Name: models.TraitsName{
				First: "first",
				Last:  "last",
			},
		},
	}
	tdata.AddIdentity(identity)

	handlers := handler.NewAuthHooksHandler(domain.NewAuthHookAPI())
	r.Post("/", handlers.HandleHookRegister)

	ut := identity.GetTraits().(*models.UserTraits)
	data := dto.HooksKratosPayloadDTO{
		UserID: identity.GetId(),
		Traits: *ut,
	}

	b, _ := json.Marshal(data)
	reader := strings.NewReader(string(b))
	req, _ := http.NewRequest("POST", "/", reader)
	req.Header.Set("Content-Type", "application/json")

	res, err := r.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	user := models.User{ID: uuid.MustParse(identity.GetId())}
	tx := database.GetDB(m.Ctx).Find(&user)
	assert.NoError(t, tx.Error)
	assert.False(t, user.CreatedAt.IsZero())
	assert.False(t, user.UpdatedAt.IsZero())
	assert.True(t, user.DeletedAt.Time.IsZero())
}
