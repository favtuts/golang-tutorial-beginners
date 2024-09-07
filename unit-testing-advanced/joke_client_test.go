package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetJokeHandler(t *testing.T) {
	table := []struct {
		id         string
		statusCode int
		body       string
	}{
		{"R7UfaahVfFd", 200, "My dog used to chase people on a bike a lot. It got so bad I had to take his bike away."},
		{"173782", 404, `Joke with id "173782" not found`},
		{"", 400, "Joke ID cannot be empty"},
	}

	for _, v := range table {
		t.Run(v.id, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "/joke?id="+v.id, nil)

			getJoke(w, r)

			if w.Code != v.statusCode {
				t.Fatalf("Expected status code: %d, but got: %d", v.statusCode, w.Code)
			}

			body := strings.TrimSpace(w.Body.String())

			if body != v.body {
				t.Fatalf("Expected body to be: '%s', but got: '%s'", v.body, body)
			}
		})
	}
}
