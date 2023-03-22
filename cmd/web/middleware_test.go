package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var th TestHandler
	h := NoSurf(&th)

	switch i := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", i)
	}
}

func TestSessionLoad(t *testing.T) {
	var th TestHandler
	h := SessionLoad(&th)

	switch i := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", i)
	}
}
