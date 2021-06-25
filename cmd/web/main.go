package main

import (
	"fmt"
	"github.com/KabirGupta07/bookings/pkg/config"
	"github.com/KabirGupta07/bookings/pkg/handlers"
	"github.com/KabirGupta07/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)
var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"
func main() {

	//Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction


	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err!=nil{
		log.Fatal(err)
	}
	app.TemplateCache = tc
	app.UseCache =false

	render.NewTemplates(&app)

	repo :=handlers.NewRepo(&app)
	handlers.NewHandlers(repo)


	srv:= http.Server{
		Addr:portNumber,
		Handler:routes(&app),
	}

	err= srv.ListenAndServe()
	fmt.Println(err)
}
