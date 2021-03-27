package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

//ref, not needed.
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

//adds csrf protect to all post requests
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

//load and saves session on every req
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}