package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/add?a=2&b=3", nil)
	rec := httptest.NewRecorder()

	AddHandler(rec, req)

	assertStatusOK(t, rec)
	assertResultEquals(t, rec, 5)
}

func TestSubtractHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/subtract?a=10&b=4", nil)
	rec := httptest.NewRecorder()

	SubtractHandler(rec, req)

	assertStatusOK(t, rec)
	assertResultEquals(t, rec, 6)
}

func TestMultiplyHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/multiply?a=6&b=7", nil)
	rec := httptest.NewRecorder()

	MultiplyHandler(rec, req)

	assertStatusOK(t, rec)
	assertResultEquals(t, rec, 42)
}

func TestDivideHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/divide?a=10&b=2", nil)
	rec := httptest.NewRecorder()

	DivideHandler(rec, req)

	assertStatusOK(t, rec)
	assertResultEquals(t, rec, 5)
}

func TestDivideByZero(t *testing.T) {
	req := httptest.NewRequest("GET", "/divide?a=10&b=0", nil)
	rec := httptest.NewRecorder()

	DivideHandler(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 for divide by zero, got %d", rec.Code)
	}
}

func TestInvalidParams(t *testing.T) {
	endpoints := []struct {
		name string
		fn   http.HandlerFunc
	}{
		{"Add", AddHandler},
		{"Subtract", SubtractHandler},
		{"Multiply", MultiplyHandler},
		{"Divide", DivideHandler},
	}

	for _, ep := range endpoints {
		t.Run(ep.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/?a=x&b=1", nil)
			rec := httptest.NewRecorder()

			ep.fn(rec, req)

			if rec.Code != http.StatusBadRequest {
				t.Errorf("expected status 400 for invalid input, got %d", rec.Code)
			}
		})
	}
}

// --- helper assertions ---

func assertStatusOK(t *testing.T, rec *httptest.ResponseRecorder) {
	t.Helper()
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func assertResultEquals(t *testing.T, rec *httptest.ResponseRecorder, expected float64) {
	t.Helper()
	var resp map[string]float64
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	result, ok := resp["result"]
	if !ok {
		t.Fatalf("missing 'result' in response")
	}
	if result != expected {
		t.Errorf("expected result %v, got %v", expected, result)
	}
}
