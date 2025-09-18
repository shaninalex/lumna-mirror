// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/org/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/org/domain"
	"gitlab.com/shaninalex/flowreon/apps/org/dto"
	"gitlab.com/shaninalex/flowreon/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_HandleGetByUser(t *testing.T) {
	m := tdata.Manager()

	org, user, _ := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)

	handlers := handler.NewOrganizationHandler(domain.NewOrganizationAPI())

	req, _ := http.NewRequest("GET", "/", nil)
	tdata.SetAuthRequest(req, user, cookie)

	rr := httptest.NewRecorder()
	dbMiddleware := database.NewMiddleware(m.DB)
	h := dbMiddleware.Wrap(http.HandlerFunc(handlers.HandleGetByUser))
	h.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	body, _ := io.ReadAll(rr.Result().Body)

	var response web.APIResponse[*dto.OrganizationDto]
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, org.Description, response.Data.Description)
	assert.Equal(t, org.Title, response.Data.Title)
}
