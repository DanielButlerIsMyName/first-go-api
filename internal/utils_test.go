package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestParseOperands_Valid(t *testing.T) {
	req := httptest.NewRequest("GET", "/?a=3.5&b=2.1", nil)
	a, b, err := ParseOperands(req)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if a != 3.5 || b != 2.1 {
		t.Errorf("expected a=3.5, b=2.1, got a=%v, b=%v", a, b)
	}
}

func TestParseOperands_InvalidA(t *testing.T) {
	req := httptest.NewRequest("GET", "/?a=foo&b=2", nil)
	_, _, err := ParseOperands(req)

	if err == nil {
		t.Error("expected error for invalid 'a', got nil")
	}
}

func TestParseOperands_InvalidB(t *testing.T) {
	req := httptest.NewRequest("GET", "/?a=2&b=bar", nil)
	_, _, err := ParseOperands(req)

	if err == nil {
		t.Error("expected error for invalid 'b', got nil")
	}
}

func TestParseOperands_MissingParams(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	_, _, err := ParseOperands(req)

	if err == nil {
		t.Error("expected error for missing parameters, got nil")
	}
}

func TestRespondJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	RespondJSON(rr, 42.0)

	// Check status code (should default to 200)
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %v", rr.Code)
	}

	// Check content type
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("expected content-type application/json, got %v", ct)
	}

	// Check response body
	var data map[string]float64
	if err := json.Unmarshal(rr.Body.Bytes(), &data); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}

	if data["result"] != 42.0 {
		t.Errorf("expected result=42.0, got %v", data["result"])
	}
}
