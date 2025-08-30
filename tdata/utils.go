// Copyright © 2025 Flowreon https://flowreon.shaninalex.com. All rights reserved.

package tdata

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// CreateTempFile - creates a new temp file.
func CreateTempFile(t *testing.T, fileName string, content []byte) string {
	filePath := filepath.Join(t.TempDir(), fileName)
	_ = os.WriteFile(filePath, content, 0644)
	fmt.Println(string(content))
	return filePath
}

// SendRequest - send request.
func SendRequest(server *httptest.Server, payload string) *http.Response {
	reader := strings.NewReader(payload)
	request, err := http.NewRequest("POST", server.URL, reader)
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	return resp
}
