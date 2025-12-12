package auth_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/pkg/token"
	"gitlab.com/shaninalex/lumna/app/services"
	"gitlab.com/shaninalex/lumna/app/tests"
	"gitlab.com/shaninalex/lumna/app/web/api/auth"
)

func Test_HandleLogin(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	userService := services.NewUserManager()
	testUser, _ := userService.CreateUser(ctx, "test@test.com", "password123")

	router := tests.NewTestRouter(ctx)
	auth.RegisterAuthController(router)
	reqBody := `{"email":"test@test.com", "password":"password123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "Should return status: \"200 OK\"")

	cookies := rr.Result().Cookies()
	cookieMap := map[string]*http.Cookie{}
	for _, c := range cookies {
		cookieMap[c.Name] = c
	}

	accessTokenManager := services.NewAccessTokenJWTService(tests.TestSecretKey, "lumna")
	refreshTokenManager := services.NewRefreshTokenJWTService(tests.TestSecretKey, "lumna")

	if assert.Contains(t, cookieMap, token.AccessTokenCookieName) {
		assert.NotEmpty(t, cookieMap[token.AccessTokenCookieName].Value)
		accessTokenClaims, err := accessTokenManager.Validate(cookieMap[token.AccessTokenCookieName].Value, token.AudTokenAPIUser)
		assert.NoError(t, err)
		id, _ := strconv.Atoi(accessTokenClaims.Subject)
		assert.Equal(t, testUser.GetId(), uint(id))
	}

	if assert.Contains(t, cookieMap, token.RefreshTokenCookieName) {
		assert.NotEmpty(t, cookieMap[token.RefreshTokenCookieName].Value)
		refreshTokenClaims, err := refreshTokenManager.Validate(cookieMap[token.AccessTokenCookieName].Value)
		assert.NoError(t, err)
		id, _ := strconv.Atoi(refreshTokenClaims.Subject)
		assert.Equal(t, testUser.GetId(), uint(id))
	}
}

func Test_HandleLoginInvalid(t *testing.T) {
	ctx := tests.TestContext()
	tests.ResetDatabase()

	userService := services.NewUserManager()
	userService.CreateUser(ctx, "test@test.com", "password123")

	router := tests.NewTestRouter(ctx)
	auth.RegisterAuthController(router)
	reqBody := `{"email":"test@test.com","password":"wrong_password"}`
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBufferString(reqBody))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusBadRequest, "Should return status: \"400 OK\"")
}
