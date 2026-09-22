package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloRoute(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}

	expected := "Hello World\n"
	got := rec.Body.String()
	if got != expected {
		t.Errorf("Expected body %q, got %q", expected, got)
	}
}
