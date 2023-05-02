package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/steffannunez/golangs/goWeb/pkg/config"
	"github.com/steffannunez/golangs/goWeb/pkg/handlers"
	"github.com/steffannunez/golangs/goWeb/pkg/renders"
)

const puerto = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//RECORDAR cambiar esto a true cuando este en prod
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := renders.CreateTemplateCacheDos()
	if err != nil {
		log.Fatal("No se pudo crear el template")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	renders.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Inicializando aplicación en el puerto %s", puerto))

	srv := &http.Server{
		Addr:    puerto,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)

}
