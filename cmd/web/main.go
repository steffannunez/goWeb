package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/steffannunez/golangs/goWeb/pkg/config"
	"github.com/steffannunez/golangs/goWeb/pkg/handlers"
	"github.com/steffannunez/golangs/goWeb/pkg/renders"
)

const puerto = ":8080"

func main() {

	var app config.AppConfig
	tc, err := renders.CreateTemplateCacheDos()
	if err != nil {
		log.Fatal("No se pudo crear el template")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)
	/*
		http.HandleFunc("/", handlers.Repo.Home)
		http.HandleFunc("/about", handlers.Repo.About)
	*/
	fmt.Println(fmt.Sprintf("Inicializando aplicación en el puerto %s", puerto))
	//_ = http.ListenAndServe(puerto, nil)

	srv := &http.Server{
		Addr:    puerto,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
