package web_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app2/pkg/db"
	"gitlab.com/shaninalex/lumna/app2/tests"
	"gitlab.com/shaninalex/lumna/app2/web"
	"gitlab.com/shaninalex/lumna/app2/web/middlewares"
)

func Test_RouterGETMethod(t *testing.T) {
	r := web.NewRouter()

	r.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	})

	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "Should return status: \"200 OK\"")
	body, _ := io.ReadAll(rr.Body)
	assert.Equal(t, string(body), "test", "Should get correct response body")
}

func Test_DefaultRouter(t *testing.T) {
	ctx := tests.TestContext()
	router := web.NewDefaultRouter()
	router.ApplyMiddlewares([]web.RouterMiddleware{
		db.NewMiddleware(db.FromContext(ctx)),
		middlewares.NewCommonMiddleware(),
		middlewares.NewHeadersMiddleware(),
	})
	router.GET("/test", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		db.FromContext(ctx)
		_, err := db.FromContext(ctx).ExecContext(ctx, `SELECT date()`)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to execute database query"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test"))
	})
	req := httptest.NewRequest("GET", "/test", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	assert.Equal(t, rr.Code, http.StatusOK, "Should return status: \"200 OK\"")
	body, _ := io.ReadAll(rr.Body)
	assert.Equal(t, string(body), "test", "Should get correct response body")
	assert.Equal(t, rr.Header().Get("X-API-Version"), "v1")
}
