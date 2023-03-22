package main

import (
	"testing"

	"github.com/garrettjhnsn/projectmf/internal/config"
	"github.com/go-chi/chi"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch i := mux.(type) {
	case *chi.Mux:
		// do nothing; test passed
	default:
		t.Errorf("type is not *chi.Mux, type is %T", i)
	}
}
