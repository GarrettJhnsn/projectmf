package handlers

import (
	"net/http"

	"github.com/garrettjhnsn/projectmf/helpers"
	"github.com/garrettjhnsn/projectmf/internal/config"
	"github.com/garrettjhnsn/projectmf/internal/driver"
	"github.com/garrettjhnsn/projectmf/internal/forms"
	"github.com/garrettjhnsn/projectmf/internal/models"
	"github.com/garrettjhnsn/projectmf/internal/render"
	"github.com/garrettjhnsn/projectmf/internal/repository"
	"github.com/garrettjhnsn/projectmf/internal/repository/dbrepo"
)

// The Repository Used By Handlers
var Repo *Repository

// The Repository Type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

// Creates A New Repository
func NewRepo(a *config.AppConfig, db *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db.SQL, a),
	}
}

// Sets The Repository For The Handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home Page Handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{})
}

// Login Page Handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "login.page.tmpl", &models.TemplateData{})
}

// Packages Page Handler
func (m *Repository) Packages(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "packages.page.tmpl", &models.TemplateData{})
}

// Consultation Page Handler
func (m *Repository) Consultation(w http.ResponseWriter, r *http.Request) {
	var emptyRequest models.Consultation
	data := make(map[string]interface{})
	data["consultationRequest"] = emptyRequest

	render.RenderTemplate(w, r, "consultation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostConsultation Handles Post Request Of Request Form
func (m *Repository) PostConsultation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	consultationRequest := models.Consultation{
		FirstName:   r.Form.Get("first_name"),
		LastName:    r.Form.Get("last_name"),
		Email:       r.Form.Get("email"),
		PhoneNumber: r.Form.Get("phone_number"),
		Date:        r.Form.Get("appointment_date"),
		Time:        r.Form.Get("appointment_time"),
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email", "phone_number", "appointment_date", "appointment_time", "check_box")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["consultationRequest"] = consultationRequest

		render.RenderTemplate(w, r, "consultation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "consultationRequest", consultationRequest)
	http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

// After Successful Submission Users Redirected To Confirmation With Entered Information Displayed Confirming Request
func (m *Repository) Confirmation(w http.ResponseWriter, r *http.Request) {
	consultationRequest, ok := m.App.Session.Get(r.Context(), "consultationRequest").(models.Consultation)

	if !ok {
		m.App.Session.Put(r.Context(), "error", "error receiving items from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	//Removing Data From Session
	m.App.Session.Remove(r.Context(), "consulationRequest")

	//Adding Data To Display
	data := make(map[string]interface{})
	data["consultationRequest"] = consultationRequest

	render.RenderTemplate(w, r, "confirmation.page.tmpl", &models.TemplateData{
		Data: data,
	})
}

// TermsOfService Page Handler
func (m *Repository) TermsOfService(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "tos.page.tmpl", &models.TemplateData{})
}

// Debug FavIcon NotFound
func (m *Repository) DoNothing(w http.ResponseWriter, r *http.Request) {}
