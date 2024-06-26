package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	cnst "github.com/LostArrows27/snippetbox/internal/const"
	"github.com/LostArrows27/snippetbox/internal/models"
	"github.com/LostArrows27/snippetbox/pkg/logger"
)

type Application struct {
	ErrorLog      logger.CustomLogger
	InfoLog       logger.CustomLogger
	Snippets      *models.SnippetModel
	TemplateCache map[string]*template.Template
}

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.Snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.html", data)

}

func (app *Application) CreateSnippetHanlder(w http.ResponseWriter, r *http.Request) {

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n–Kobayashi Issa"
	expires := 7
	id, err := app.Snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id),
		http.StatusSeeOther)

}

func (app *Application) ViewSnippetHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)

	if err != nil || id < 0 {
		app.notFound(w)
		return
	}

	snippet, err := app.Snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.html", data)

}

func (app *Application) StaticFileHanlder(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir(cnst.StaticFileDir))

	handlerFile := http.StripPrefix("/static", fileServer)

	handlerFile.ServeHTTP(w, r)
}
