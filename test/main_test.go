package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// TestHandler checks if the HTTP handler responds correctly.
func TestHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }
    rec := httptest.NewRecorder()
    handler(rec, req)

    if rec.Code != http.StatusOK {
        t.Errorf("Expected status OK; got %v", rec.Code)
    }

    expected := "Hello World!"
    if rec.Body.String() != expected {
        t.Errorf("Expected body '%s'; got '%s'", expected, rec.Body.String())
    }
}
