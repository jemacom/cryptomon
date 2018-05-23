package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPricingHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/v1/pricing", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(pricingHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
