package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRankingHandler(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"Without limit param", "/v1/ranking"},
		{"limit is 0", "/v1/ranking?limit=0"},
		{"limit is negative", "/v1/ranking?limit=-2"},
		{"limit is 4", "/v1/ranking?limit=4"},
	}

	for _, test := range tests {
		req, err := http.NewRequest("GET", test.url, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := http.HandlerFunc(rankingHandler)

		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	}

}
