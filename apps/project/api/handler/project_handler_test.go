// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/project/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_ProjectList(t *testing.T) {
	m := tdata.Manager()
	tdata.Clear(m.Ctx)

	_, user, project := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)

	r := tdata.AuthTestRouter()
	handlers := handler.NewProjectHandler(domain.NewProjectManagement())
	r.Get("/", handlers.HandleProjectsList)

	req, _ := http.NewRequest("GET", "/", nil)
	tdata.SetAuthRequest(req, user, cookie)

	res, err := r.Test(req, -1)

	assert.Nil(t, err)
	assert.Equal(t, 200, res.StatusCode)

	body, _ := io.ReadAll(res.Body)

	var response web.APIResponse[[]*dto.ProjectDto]
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, project.GetID(), response.Data[0].ID)
}

func Test_ProjectTaskList(t *testing.T) {
	m := tdata.Manager()
	tdata.Clear(m.Ctx)

	_, user, project := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)

	tasks := tdata.CreateTasks(m.Ctx, project)

	r := tdata.AuthTestRouter()
	handlers := handler.NewProjectHandler(domain.NewProjectManagement())
	r.Get("/api/project/:projectCode/tasks", handlers.HandleProjectTasksList)

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/project/%s/tasks", project.ProjectKey), nil)
	tdata.SetAuthRequest(req, user, cookie)

	res, err := r.Test(req, -1)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	body, _ := io.ReadAll(res.Body)

	var response web.APIResponse[[]*dto.TaskDto]
	err = json.Unmarshal(body, &response)
	assert.NoError(t, err)
	assert.Equal(t, len(tasks), len(response.Data))
}
