package handlers

import (
	"net/http"

	"github.com/SergioPopovs176/booking/pkg/config"
	"github.com/SergioPopovs176/booking/pkg/models"
	"github.com/SergioPopovs176/booking/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	// fmt.Println(r.Context())

	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// send the data to the template

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}

func (m *Repository) Links(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "links.page.tmpl", &models.TemplateData{})
}
