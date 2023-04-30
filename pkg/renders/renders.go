package renders

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/steffannunez/golangs/goWeb/pkg/config"
	"github.com/steffannunez/golangs/goWeb/pkg/models"
)

var funciones = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCacheDos()
	}
	//crear un cache para templates

	//get requested template desde el cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("No pude crear templates desde templates cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	// render del template

	_, err := buf.WriteTo(w)

	if err != nil {
		log.Println("Error escribiendo el template al browser", err)
	}

}

var tc = make(map[string]*template.Template)

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// verificar si podermos guardar el template en cache

	_, inMap := tc[t]
	if !inMap {
		//tenemos que crear un template
		log.Println("creando teemplate y añadiendola al cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//tenemos el templatee en cache
		log.Println("usando template en cache")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	//parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache map

	tc[t] = tmpl
	return nil
}

func CreateTemplateCacheDos() (map[string]*template.Template, error) {
	miCache := map[string]*template.Template{} //esto es igual que usar el make()
	//obtener todos los files llamados *page.tmpl en la carpeta ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return miCache, err
	}

	//range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return miCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if err != nil {
			return miCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return miCache, err
			}
		}
		miCache[name] = ts
	}
	return miCache, nil
}
