package client

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUploadCase(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("expected POST, got %s", r.Method)
			}
			if ct := r.Header.Get("Content-Type"); ct != "application/zip" {
				t.Errorf("expected Content-Type application/zip, got %s", ct)
			}
			body, _ := io.ReadAll(r.Body)
			if string(body) != "fake-zip-bytes" {
				t.Errorf("unexpected body: %s", string(body))
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"id":"abc-123","status":"processing"}`))
		}))
		defer server.Close()

		result, err := UploadCase(server.URL, []byte("fake-zip-bytes"))
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "abc-123" {
			t.Errorf("expected id abc-123, got %s", result.ID)
		}
		if result.Status != "processing" {
			t.Errorf("expected status processing, got %s", result.Status)
		}
	})

	t.Run("backend error 400", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"bad zip"}`))
		}))
		defer server.Close()

		_, err := UploadCase(server.URL, []byte("fake"))
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "400") {
			t.Errorf("expected error to mention 400, got %v", err)
		}
	})
}

func TestGetStatus(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/cases/abc-123/status" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":"abc-123","status":"done"}`))
		}))
		defer server.Close()

		result, err := GetStatus(server.URL, "abc-123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result.ID != "abc-123" {
			t.Errorf("expected id abc-123, got %s", result.ID)
		}
		if result.Status != "done" {
			t.Errorf("expected status done, got %s", result.Status)
		}
	})

	t.Run("not found 404", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"not found"}`))
		}))
		defer server.Close()

		_, err := GetStatus(server.URL, "missing")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "404") {
			t.Errorf("expected error to mention 404, got %v", err)
		}
	})
}

func TestGetCase(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				t.Errorf("expected GET, got %s", r.Method)
			}
			if r.URL.Path != "/cases/abc-123" {
				t.Errorf("unexpected path: %s", r.URL.Path)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"id":"abc-123","status":"done","alert":{"signature":"test"}}`))
		}))
		defer server.Close()

		result, err := GetCase(server.URL, "abc-123")
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		c := *result
		if c["id"] != "abc-123" {
			t.Errorf("expected id abc-123, got %v", c["id"])
		}
		if c["status"] != "done" {
			t.Errorf("expected status done, got %v", c["status"])
		}
	})

	t.Run("not found 404", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"not found"}`))
		}))
		defer server.Close()

		_, err := GetCase(server.URL, "missing")
		if err == nil {
			t.Fatal("expected error, got nil")
		}
		if !strings.Contains(err.Error(), "404") {
			t.Errorf("expected error to mention 404, got %v", err)
		}
	})
}
