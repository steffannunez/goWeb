package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

// WruteToConsole sirve para poner en la consola lo que pase en la pagina
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Recargando la página")
		next.ServeHTTP(w, r)

	})
}

// NoSurf adds CSRF protection to all POST requests
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

// SessionLoad carga y guarda la sesion en todas las requests
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
