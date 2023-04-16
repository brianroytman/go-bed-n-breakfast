package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/brianroytman/go-bnb-course/pkg/config"
	"github.com/brianroytman/go-bnb-course/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates - sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate - renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var templateCache map[string]*template.Template
	// get the template cache from the app config
	if app.UseCache {
		templateCache = app.TemplateCache

	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
