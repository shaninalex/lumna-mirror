// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package handler_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/flowreon/apps/project/adapter"
	"gitlab.com/shaninalex/flowreon/apps/project/api/handler"
	"gitlab.com/shaninalex/flowreon/apps/project/domain"
	"gitlab.com/shaninalex/flowreon/apps/project/dto"
	"gitlab.com/shaninalex/flowreon/internal/database"
	"gitlab.com/shaninalex/flowreon/internal/web"
	"gitlab.com/shaninalex/flowreon/models"
	"gitlab.com/shaninalex/flowreon/tdata"
)

func Test_HandleTaskPatchStatus(t *testing.T) {
	m := tdata.Manager()

	_, user, project := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)
	statuses := tdata.CreateRandomTasksAndStatuses(m.Ctx, project)
	handlers := handler.NewTaskHandler(domain.NewProjectManagement())

	router := tdata.AuthTestRouter(m.Ctx)
	router.PATCH("/api/project/{projectCode}/tasks/{taskCode}/status", handlers.HandleTaskPatchStatus)

	testTask := statuses[0].Tasks[0]

	payload := dto.ChangeTaskStatusDTO{
		FromStatusID: testTask.TaskStatusID,
		ToStatusID:   statuses[1].ID,
		FromIdx:      1,
		ToIdx:        0,
	}

	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/project/%s/tasks/%s/status", project.ProjectKey, testTask.Code), strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")
	tdata.SetAuthRequest(req, user, cookie)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Result().StatusCode)
	body, _ := io.ReadAll(rr.Result().Body)
	assert.NotNil(t, body)
	var response web.APIResponse[any]
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)
	assert.Equal(t, []string{"Task saved"}, response.Messages)

	taskFromDB := models.Task{ID: testTask.ID}
	database.GetDB(m.Ctx).First(&taskFromDB)
	assert.Equal(t, statuses[1].ID, taskFromDB.TaskStatusID)
}

func Test_HandleTaskDetail(t *testing.T) {
	m := tdata.Manager()

	_, user, project := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)
	tasks := tdata.CreateTasks(m.Ctx, project)
	handlers := handler.NewTaskHandler(domain.NewProjectManagement())
	router := tdata.AuthTestRouter(m.Ctx)
	router.GET("/api/project/{projectCode}/tasks/{taskCode}", handlers.HandleTaskDetail)

	testTask := tasks[0]

	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/project/%s/tasks/%s", project.ProjectKey, testTask.Code), nil)
	tdata.SetAuthRequest(req, user, cookie)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Result().StatusCode)

	body, _ := io.ReadAll(rr.Result().Body)

	var response web.APIResponse[*dto.TaskDto]
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)

	testTaskDto := adapter.NewTaskDto(testTask)
	assert.Equal(t, testTaskDto.ID, response.Data.ID)
	assert.Equal(t, testTaskDto.UserID, response.Data.UserID)
	assert.Equal(t, testTaskDto.EpicID, response.Data.EpicID)
	assert.Equal(t, testTaskDto.SprintID, response.Data.SprintID)
	assert.Equal(t, testTaskDto.ProjectID, response.Data.ProjectID)
	assert.Equal(t, testTaskDto.Assignee, response.Data.Assignee)
	assert.Equal(t, testTaskDto.Completed, response.Data.Completed)
	assert.Equal(t, testTaskDto.Title, response.Data.Title)
	assert.Equal(t, testTaskDto.Description, response.Data.Description)
	assert.Equal(t, testTaskDto.StatusID, response.Data.StatusID)
	assert.Equal(t, testTaskDto.ListIdx, response.Data.ListIdx)
	assert.Equal(t, testTaskDto.Code, response.Data.Code)
	assert.WithinDuration(t, testTaskDto.CreatedAt, response.Data.CreatedAt, time.Millisecond)
	assert.WithinDuration(t, testTaskDto.UpdatedAt, response.Data.UpdatedAt, time.Millisecond)
}

func Test_HandleTaskUpdate(t *testing.T) {
	m := tdata.Manager()

	_, user, project := tdata.CreatePack(m.Ctx)
	cookie := tdata.AddSession(user)
	statuses := tdata.CreateRandomTasksAndStatuses(m.Ctx, project)
	handlers := handler.NewTaskHandler(domain.NewProjectManagement())
	router := tdata.AuthTestRouter(m.Ctx)
	router.PATCH("/api/project/{projectCode}/tasks/{taskCode}", handlers.HandleTaskUpdate)

	testTask := statuses[0].Tasks[0]
	payload := adapter.UpdateTaskData{
		Title:       uuid.NewString(),
		Description: uuid.NewString(),
	}

	b, _ := json.Marshal(payload)
	req, _ := http.NewRequest("PATCH", fmt.Sprintf("/api/project/%s/tasks/%s", project.ProjectKey, testTask.Code), strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")
	tdata.SetAuthRequest(req, user, cookie)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Result().StatusCode)
	body, _ := io.ReadAll(rr.Result().Body)
	assert.NotNil(t, body)
	var response web.APIResponse[any]
	err := json.Unmarshal(body, &response)
	assert.NoError(t, err)
	assert.Equal(t, []string{"Task saved"}, response.Messages)

	taskFromDB := models.Task{ID: testTask.ID}
	database.GetDB(m.Ctx).First(&taskFromDB)

	assert.Equal(t, payload.Title, taskFromDB.Title)
	assert.Equal(t, payload.Description, taskFromDB.Description)
}
