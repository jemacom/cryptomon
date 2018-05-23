package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCoinsHandler(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"Without limit param", "/v1"},
		{"limit is 0", "/v1?limit=0"},
		{"limit is negative", "/v1?limit=-2"},
		{"limit is 4", "/v1?limit=4"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			listCoinsHandler(w, r, "http://localhost:8081/v1/pricing", "http://localhost:8082/v2/ranking")
		})

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}
}
