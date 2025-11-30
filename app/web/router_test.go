package web_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/web"
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
	t.Log(rr.Body)
	assert.Equal(t, string(body), "test", "Should get correct response body")
}
