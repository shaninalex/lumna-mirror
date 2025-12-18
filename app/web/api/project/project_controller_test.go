package project_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/db"
	"gitlab.com/shaninalex/lumna/app/repositories"
	"gitlab.com/shaninalex/lumna/app/tests"
	"gitlab.com/shaninalex/lumna/app/web/api/project"
)

func Test_ApiProjectController_List(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()
	router := tests.NewTestRouter(ctx)

	p := &models.Project{Name: "test"}
	repo := repositories.NewProjectRespository()
	repo.Create(ctx, p)

	project.RegisterProjectController(router)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/projects", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "Should return status: \"200 OK\"")

	b, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)

	assert.Contains(t, string(b), p.Name, "Should contain project name")
	assert.Contains(t, string(b), fmt.Sprintf("%d", p.GetId()), "Should contain project id")
}

func TTest_ApiProjectController_Create(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	repo := repositories.NewProjectRespository()

	router := tests.NewTestRouter(ctx)
	project.RegisterProjectController(router)

	reqBody := `{"name":"test"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/projects", bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")

	b, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(b), "test", "Should contain project name")

	projectsList, err := repo.List(ctx, db.Eq("name", "test"))
	assert.NoError(t, err)
	assert.Equal(t, 1, len(projectsList))
	assert.Equal(t, "test", projectsList[0].Name)
}

func Test_ApiProjectController_BoardsList(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	p := tests.CreateProjectWithName(ctx, "test")
	boardA := tests.CreateBoard(ctx, p.GetId(), "board A")
	boardB := tests.CreateBoard(ctx, p.GetId(), "board B")
	router := tests.NewTestRouter(ctx)
	project.RegisterProjectController(router)
	url := fmt.Sprintf("/api/v1/project/%d/boards", p.GetId())
	req := httptest.NewRequest(http.MethodGet, url, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")

	b, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(b), boardA.Name, "Should contain board name")
	assert.Contains(t, string(b), boardB.Name, "Should contain board name")
}

func Test_ApiProjectController_BoardsCreate(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	p := tests.CreateProjectWithName(ctx, "test")
	router := tests.NewTestRouter(ctx)
	project.RegisterProjectController(router)

	url := fmt.Sprintf("/api/v1/project/%d/boards", p.GetId())
	boardName := "development"
	reqBody := fmt.Sprintf("{\"name\": \"%s\"}", boardName)
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "Should return status: \"200 OK\"")

	b, err := io.ReadAll(rr.Body)
	assert.NoError(t, err)
	assert.Contains(t, string(b), boardName, "Should contain board name")
}
