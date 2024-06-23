package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"

	cnst "github.com/LostArrows27/snippetbox/internal/const"
	"github.com/LostArrows27/snippetbox/internal/models"
)

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}

/*
in this function
1. We get all the page template from html/pages folder
2. Create template with that page
3. Save in a map with format:
  - key -> page name
  - value -> *template.Template

EX. template["home.html"]
*/
func NewTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob(cnst.PagesFileSearchPattern)
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.ParseFiles(cnst.HomeBasePath)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(cnst.PartialsFileSearchPattern)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

// use to render page based on page name with template data
// EX. home.html / view.html
func (app *Application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.TemplateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
