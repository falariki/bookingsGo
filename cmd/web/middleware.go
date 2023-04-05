package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protecyion to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

//loading the session
// LoadAndSave provides middleware which automatically loads and saves session
// data for the current request, and communicates the session token to and from
// the client in a cookie

// SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
