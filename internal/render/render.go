package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/garrettjhnsn/projectmf/internal/config"
	"github.com/garrettjhnsn/projectmf/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// Sets The Config For The Template Package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData, r *http.Request) {
	var tc map[string]*template.Template

	// Get The Template Cache From App Config
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// Get Requested Template From Cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get requested template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get All Files Name *.page.tmpl From ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Range Through All Files With *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Looking For Layout Files
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// INTERMEDIATE --> Renders Template Using html/template with caching

/*
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var err error
	// Check To See If We Already Have The Template In Our Cache
	_, inMap := tc[t]

	if !inMap {
		// Need To Create Template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		// We Have The Template In Cache
		log.Println("using cached template")
	}

	tmpl := tc[t]

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

	// Parse The Template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// Add Template To Cache (Map)
	tc[t] = tmpl
ResponseWriter
	return nil
}
*/

// BASIC --> Renders Template Using html/template //

/*
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error parsing template:", err)
	}
}
*/
