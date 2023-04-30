package handlers

import (
	"net/http"

	"github.com/steffannunez/golangs/goWeb/pkg/config"
	"github.com/steffannunez/golangs/goWeb/pkg/models"
	"github.com/steffannunez/golangs/goWeb/pkg/renders"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository typee
type Repository struct {
	App *config.AppConfig
}

// New Repo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for de handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	renders.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// logica
	stringMap := make(map[string]string)
	stringMap["test"] = "Hola putines!"
	renders.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
