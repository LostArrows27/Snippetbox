package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"unicode/utf8"

	cnst "github.com/LostArrows27/snippetbox/internal/const"
	"github.com/LostArrows27/snippetbox/internal/models"
	"github.com/LostArrows27/snippetbox/pkg/logger"
	"github.com/julienschmidt/httprouter"
)

type Application struct {
	ErrorLog      logger.CustomLogger
	InfoLog       logger.CustomLogger
	Snippets      *models.SnippetModel
	TemplateCache map[string]*template.Template
}

func (app *Application) snippetHomeView(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.Snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, http.StatusOK, "home.html", data)

}

func (app *Application) snippetCreateForm(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "create.html", data)
}

func (app *Application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	// 1. parse + get form body data
	err := r.ParseForm()

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// 2. validate form data
	fieldErrors := make(map[string]string)

	if strings.TrimSpace(title) == "" {
		fieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		fieldErrors["title"] = "This field cannot be more than 100 characters long"
	}

	if strings.TrimSpace(content) == "" {
		fieldErrors["content"] = "This field cannot be blank"
	}

	if expires != 1 && expires != 7 && expires != 365 {
		fieldErrors["expires"] = "This field must equal 1, 7 or 365"
	}

	if len(fieldErrors) > 0 {
		fmt.Fprint(w, fieldErrors)
		return
	}

	// 3. insert form to database + response
	id, err := app.Snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id),
		http.StatusSeeOther)

}

func (app *Application) snippetView(w http.ResponseWriter, r *http.Request) {

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil || id < 1 {
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

func (app *Application) fileHandler(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir(cnst.StaticFileDir))

	handlerFile := http.StripPrefix("/static", fileServer)

	handlerFile.ServeHTTP(w, r)
}
