package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	index(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status: %d, but got: %d", http.StatusOK, w.Code)
	}
}
