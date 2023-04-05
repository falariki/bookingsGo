package main

import (
	"net/http"

	"github.com/falariki/bookingsProject/pkg/config"
	"github.com/falariki/bookingsProject/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	//using middleware

	mux.Use(middleware.Recoverer)
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}

//every time sombody hits the page - print something to the console

//Sessions
// we send a request to the server
// it formats a response and sends it back to the client
// web server is stateless - it doesn't remember anything about the client
// choose
