package task_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app2/services"
	"gitlab.com/shaninalex/lumna/app2/tests"
	"gitlab.com/shaninalex/lumna/app2/web/api/task"
)

func Test_ApiTaskController_Get(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	pr := tests.CreateProject(ctx)
	br := tests.CreateBoard(ctx, pr.GetId(), "A")
	ls := tests.CreateBoardList(ctx, pr.GetId(), "A")
	tt := tests.CreateTask(ctx, br.GetId(), ls.GetId(), "A")

	task.RegisterTaskController(router)

	url := fmt.Sprintf("/api/v1/task/%d", tt.GetId())

	req := httptest.NewRequest(http.MethodGet, url, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")
	assert.Contains(t, rr.Body.String(), tt.Name)
}

func Test_ApiTaskController_Patch(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	pr := tests.CreateProject(ctx)
	br := tests.CreateBoard(ctx, pr.GetId(), "A")
	ls := tests.CreateBoardList(ctx, pr.GetId(), "A")
	tt := tests.CreateTask(ctx, br.GetId(), ls.GetId(), "A")

	task.RegisterTaskController(router)

	newName := "new name"
	url := fmt.Sprintf("/api/v1/task/%d", tt.GetId())
	reqBody := fmt.Sprintf("{\"name\":\"%s\", \"list_id\": %d}", newName, tt.ListId)

	req := httptest.NewRequest(http.MethodPatch, url, bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")

	service := services.NewTaskManager()
	dbTask, err := service.Get(ctx, tt.GetId())
	assert.NoError(t, err)
	assert.Equal(t, newName, dbTask.Name)
	assert.Equal(t, tt.GetId(), dbTask.GetId())
}

func Test_ApiTaskController_Delete(t *testing.T) {
	ctx := tests.Context()
	router := tests.NewTestRouter(ctx)

	pr := tests.CreateProject(ctx)
	br := tests.CreateBoard(ctx, pr.GetId(), "A")
	ls := tests.CreateBoardList(ctx, pr.GetId(), "A")
	tt := tests.CreateTask(ctx, br.GetId(), ls.GetId(), "A")

	task.RegisterTaskController(router)

	url := fmt.Sprintf("/api/v1/task/%d", tt.GetId())

	req := httptest.NewRequest(http.MethodDelete, url, nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")
	assert.Contains(t, rr.Body.String(), "Task deleted")
}
