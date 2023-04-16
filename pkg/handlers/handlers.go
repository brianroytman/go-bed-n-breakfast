package handlers

import (
	"net/http"

	"github.com/brianroytman/go-bnb-course/pkg/config"
	"github.com/brianroytman/go-bnb-course/pkg/models"
	"github.com/brianroytman/go-bnb-course/pkg/render"
)

// Repo - the repository used by the handlers
var Repo *Repository

// Repository - is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo - create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers - sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home - home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About - the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
