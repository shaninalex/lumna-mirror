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

func Test_WebProjectList(t *testing.T) {
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

func Test_WebProjectCreate(t *testing.T) {
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

	projectsList, err := repo.List(ctx, db.Option{Key: "name", Value: "test"})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(projectsList))
	assert.Equal(t, "test", projectsList[0].Name)
}
