package apiserver

import (
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_HandleHello(t *testing.T) {
	s := NewAPIServer(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
