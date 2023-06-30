package handlers

import (
	"golang-webapp/package/config"
	"golang-webapp/package/render"
	"net/http"
)

var Repo *Repository

// Repository is type of repository
type Repository struct {
	App *config.AppConfig
}

// NewRepo create new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets repository handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
