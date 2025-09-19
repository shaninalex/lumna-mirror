// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/project/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_ProjectList(t *testing.T) {
	m := tdata.Manager()

	_, user, project := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)

	handlers := handler.NewProjectHandler(domain.NewProjectManagement())
	router := tdata.AuthTestRouter(m.Ctx)
	router.GET("/", handlers.HandleProjectsList)

	req, _ := http.NewRequest("GET", "/", nil)
	tdata.SetAuthRequest(req, user, cookie)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Result().StatusCode)

	body, _ := io.ReadAll(rr.Result().Body)

	var response web.APIResponse[[]*dto.ProjectDto]
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, project.GetID(), response.Data[0].ID)
}

func Test_ProjectTaskList(t *testing.T) {
	m := tdata.Manager()

	_, user, project := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)

	tasks := tdata.CreateTasks(m.Ctx, project)

	handlers := handler.NewProjectHandler(domain.NewProjectManagement())
	router := tdata.AuthTestRouter(m.Ctx)
	router.GET("/api/project/{projectCode}/tasks", handlers.HandleProjectTasksList)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/project/%s/tasks", project.ProjectKey), nil)
	tdata.SetAuthRequest(req, user, cookie)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)

	body, _ := io.ReadAll(rr.Result().Body)

	var response web.APIResponse[[]*dto.TaskDto]
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)
	assert.Equal(t, len(tasks), len(response.Data))
}

func Test_ProjectCreate(t *testing.T) {
	m := tdata.Manager()
	_, user, _ := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)

	handlers := handler.NewProjectHandler(domain.NewProjectManagement())

	router := tdata.AuthTestRouter(m.Ctx)
	router.POST("/", handlers.HandleProjectCreate)
	title := uuid.NewString()
	projectDto := &dto.ProjectDto{
		Title: title,
	}
	b, _ := json.Marshal(projectDto)
	req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(string(b)))
	req.Header.Set("Content-Type", "application/json")
	tdata.SetAuthRequest(req, user, cookie)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Result().StatusCode)

	body, _ := io.ReadAll(rr.Result().Body)

	var response web.APIResponse[*dto.ProjectDto]
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)

	assert.Equal(t, title, response.Data.Title)
}
