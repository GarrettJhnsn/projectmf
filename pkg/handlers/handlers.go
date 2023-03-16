package handlers

import (
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
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// MyIp Page Handler
func (m *Repository) MyIp(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "myip.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Login Page Handler
func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "login.page.tmpl", &models.TemplateData{})
}

// Debug FavIcon NotFound
func (m *Repository) DoNothing(w http.ResponseWriter, r *http.Request) {}
