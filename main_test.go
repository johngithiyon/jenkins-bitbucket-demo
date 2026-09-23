package main

import (
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	hello(rec, req)

	expected := "Hello from Jenkins + Bitbucket!\n"

	if rec.Body.String() != expected {
		t.Errorf("expected %q, got %q", expected, rec.Body.String())
	}
}
