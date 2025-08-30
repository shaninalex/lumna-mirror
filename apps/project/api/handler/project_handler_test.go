// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler_test

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/project/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func TestProjectList(t *testing.T) {
	m := tdata.Manager()
	testRouter := web.DefaultRouter(m.DB, "test_project_api")

	handlers := handler.NewProjectHandler(domain.NewProjectManagement())
	testRouter.Get("/", handlers.HandleProjectsList)

	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/", nil)

	// Perform the request using app.Test
	res, err := testRouter.Test(req, -1)

	// Verify that no error occurred
	assert.Nil(t, err)

	// Verify the status code
	assert.Equal(t, 200, res.StatusCode)

	// Read the response body
	body, _ := io.ReadAll(res.Body)

	// Verify the response body
	assert.Equal(t, "OK", string(body))
}
