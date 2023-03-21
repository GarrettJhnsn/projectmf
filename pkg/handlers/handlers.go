package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/garrettjhnsn/projectmf/pkg/config"
	"github.com/garrettjhnsn/projectmf/pkg/models"
	"github.com/garrettjhnsn/projectmf/pkg/render"
)

// The Repository Used By Handlers
var Repo *Repository

// The Repository Type
type Repository struct {
	App *config.AppConfig
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// Creates A New Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets The Repository For The Handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{}, r)
}

// Success Page Handler
func (m *Repository) Success(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "success.page.tmpl", &models.TemplateData{}, r)
}

// Login Page Handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{}, r)
}

// Packages Page Handler
func (m *Repository) Packages(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "packages.page.tmpl", &models.TemplateData{}, r)
}

// Consultation Page Handler
func (m *Repository) Consultation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "consultation.page.tmpl", &models.TemplateData{}, r)
}

// PostConsultation Handles Post Request
func (m *Repository) PostConsultation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "consultation.page.tmpl", &models.TemplateData{}, r)
}

// ConsultationJSON Handles Request For Availability And Sends JSON Response
func (m *Repository) ConsultationJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      false,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "	")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// TermsOfService Page Handler
func (m *Repository) TermsOfService(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "tos.page.tmpl", &models.TemplateData{}, r)
}

// Debug FavIcon NotFound
func (m *Repository) DoNothing(w http.ResponseWriter, r *http.Request) {}
