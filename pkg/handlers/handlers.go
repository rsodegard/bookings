package handlers

import (
	"net/http"

	"github.com/rsodegard/bookings/pkg/config"
	"github.com/rsodegard/bookings/pkg/models"
	"github.com/rsodegard/bookings/pkg/render"
)

// the repository used by the handlers
var Repo *Repository

//is the repository type
type Repository struct {
	App *config.AppConfig
}

//creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler, with access to repository
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//About is the about page handler, with access to repository
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, Again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
