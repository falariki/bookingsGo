package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/falariki/bookingsProject/pkg/config"
	"github.com/falariki/bookingsProject/pkg/handlers"
	"github.com/falariki/bookingsProject/pkg/render"

	"github.com/alexedwards/scs/v2"
)

//different between const and var
//var just a variable - so can be changed
//const can't be changed - as it is a constant
const portNumber = ":8092"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false

	// session will last for 24 hours
	session = scs.New()
	session.Lifetime = 24 * time.Hour

	//set parameters for cookie
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannt create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
