package handlers

import (
	"fmt"
	"github.com/KabirGupta07/bookings/pkg/config"
	"github.com/KabirGupta07/bookings/pkg/models"
	"github.com/KabirGupta07/bookings/pkg/render"
	"net/http"
)



//Repo is the repository used by the handlers
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig)*Repository  {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for handlers
func NewHandlers(r *Repository)  {
	Repo = r
}

//Home is the home page handler
func(m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	fmt.Println(remoteIP)
	m.App.Session.Put(r.Context(),"remote_ip", remoteIP)

	stringMap := make(map[string]string)
	stringMap["test1"] = "Hello, Home"

	//send the data to the template
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

//About is the home page handler
func(m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again!"

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap ["remote_ip"] = remoteIP

	//send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})


}
