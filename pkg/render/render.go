package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/falariki/bookingsProject/pkg/config"
	"github.com/falariki/bookingsProject/pkg/models"
)

var app *config.AppConfig

//NewTemplates sets the config for the template package
func NewTemplates(application *config.AppConfig) {
	app = application
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {

	return templateData
}

// RenderTempletes renders templates using html/template
func RenderTempletes(w http.ResponseWriter, html string, templateData *models.TemplateData) {
	//get the template cache from the app config - no need to create every time
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := templateCache[html]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	//will hold bytes
	buffer := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)

	//ignore the result and just execute the template
	_ = t.Execute(buffer, templateData)

	//render and parse the template
	_, err := buffer.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//create a map to store the templates in - empty map
	myCache := map[string]*template.Template{}

	//get all of the files names *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	//loop through all files ending with *page.html
	for _, page := range pages {
		name := filepath.Base(page)

		//parse the file
		tampleteSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//look for layout files
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			tampleteSet, err = tampleteSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		//key name - template set as value for that key
		myCache[name] = tampleteSet
	}

	return myCache, nil
}
