package main

import (
	"net/http"

	"github.com/garrettjhnsn/projectmf/internal/config"
	"github.com/garrettjhnsn/projectmf/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/favicon", handlers.Repo.DoNothing)
	mux.Get("/success", handlers.Repo.Success)
	mux.Get("/login", handlers.Repo.Login)
	mux.Get("/packages", handlers.Repo.Packages)

	mux.Get("/consultation", handlers.Repo.Consultation)
	mux.Post("/consultation", handlers.Repo.PostConsultation)
	mux.Get("/consultation-available", handlers.Repo.ConsultationJSON)

	mux.Get("/tos", handlers.Repo.TermsOfService)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
