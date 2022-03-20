package handlers

import (
	"net/http"

	"github.com/brianweber2/bookings/pkg/config"
	"github.com/brianweber2/bookings/pkg/models"
	"github.com/brianweber2/bookings/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// GeneralsQuarters is the generals-quarters page handler
func (m *Repository) GeneralsQuarters(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals.html", &models.TemplateData{})
}

// MajorsSuite is the majors-suite page handler
func (m *Repository) MajorsSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors.html", &models.TemplateData{})
}

// Contact is the contact page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.html", &models.TemplateData{})
}

// MakeReservation is the make reservation page handler
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.html", &models.TemplateData{})
}
