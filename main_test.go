package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	health(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rr.Code)
	}
	if got := strings.TrimSpace(rr.Body.String()); got != "ok" {
		t.Fatalf("expected body 'ok', got %q", got)
	}
}

func TestHello(t *testing.T) {
	t.Run("default name", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/hello", nil)
		rr := httptest.NewRecorder()

		hello(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
			t.Fatalf("expected Content-Type application/json, got %q", ct)
		}
		var payload map[string]string
		if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}
		if payload["message"] != "hello world" {
			t.Fatalf("expected message 'hello world', got %q", payload["message"])
		}
	})

	t.Run("with name", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/hello?name=David", nil)
		rr := httptest.NewRecorder()

		hello(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d", rr.Code)
		}
		var payload map[string]string
		if err := json.Unmarshal(rr.Body.Bytes(), &payload); err != nil {
			t.Fatalf("invalid JSON: %v", err)
		}
		if payload["message"] != "hello David" {
			t.Fatalf("expected message 'hello David', got %q", payload["message"])
		}
	})
}

