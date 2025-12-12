package auth_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/tests"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
)

func Test_HandleRegister(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	router := tests.NewTestRouter(ctx)
	auth.RegisterAuthController(router)
	reqBody := `{"email":"test@test.com", "password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "Should return status: \"200 OK\"")

	userService := services.NewUserManager()
	newUser, err := userService.GetUserByEmail(ctx, "test@test.com")
	assert.NoError(t, err)
	assert.NotNil(t, newUser)
}
