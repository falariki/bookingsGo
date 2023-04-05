package handlers

import (
	"net/http"

	"github.com/falariki/bookingsProject/pkg/config"
	"github.com/falariki/bookingsProject/pkg/models"
	"github.com/falariki/bookingsProject/pkg/render"
)

//repository pattern
// allows us to swap components
//with the minimal amount of code change

//Repo type of Repository the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	//app with the type of config.AppConfig
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	//doesn't return anything just sets the variable
	Repo = repo
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//grab remote ip address and store it

	remoteIP := r.RemoteAddr
	//put that into session

	//every time somebody hit this page
	//we're going to store remote IP in the session
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTempletes(w, "home.page.html", &models.TemplateData{})
}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)

	stringMap["LFG"] = "Lets Fucking GO"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	render.RenderTempletes(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
