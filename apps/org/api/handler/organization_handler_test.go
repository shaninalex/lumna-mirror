// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/org/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/org/domain"
	"gitlab.com/shaninalex/flowreon/apps/org/dto"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_HandleGetByUser(t *testing.T) {
	m := tdata.Manager()
	tdata.Clear(m.Ctx)

	org, user, _ := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)

	r := tdata.AuthTestRouter()
	handlers := handler.NewOrganizationHandler(domain.NewOrganizationApi())
	r.Get("/", handlers.HandleGetByUser)
	req, _ := http.NewRequest("GET", "/", nil)
	tdata.SetAuthRequest(req, user, cookie)

	res, err := r.Test(req, -1)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)
	body, _ := io.ReadAll(res.Body)

	var response web.ApiResponse[*dto.OrganizationDto]
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, org.Description, response.Data.Description)
	assert.Equal(t, org.Title, response.Data.Title)
}
