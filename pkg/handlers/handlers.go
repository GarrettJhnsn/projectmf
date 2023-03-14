package handlers

import (
	"net/http"

	"github.com/garrettjhnsn/projectmf/pkg/render"
)

// Home Page Handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About Page Handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
