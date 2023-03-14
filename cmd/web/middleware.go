package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// Adds CSRF Protection To All Post Request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// Loads And Saves The Session On Every Request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
