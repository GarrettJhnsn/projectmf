package main

import (
	"net/http"

	"github.com/garrettjhnsn/projectmf/pkg/config"
	"github.com/garrettjhnsn/projectmf/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/favicon", handlers.Repo.DoNothing)

	return mux
}
