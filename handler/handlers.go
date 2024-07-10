package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	cnst "github.com/LostArrows27/snippetbox/internal/const"
	"github.com/LostArrows27/snippetbox/internal/models"
	"github.com/LostArrows27/snippetbox/internal/validator"
	"github.com/LostArrows27/snippetbox/pkg/logger"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	"github.com/julienschmidt/httprouter"
)

type Application struct {
	ErrorLog       logger.CustomLogger
	InfoLog        logger.CustomLogger
	Snippets       *models.SnippetModel
	TemplateCache  map[string]*template.Template
	FormDecoder    *form.Decoder
	SessionManager *scs.SessionManager
}

type snippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
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
	data.Form = snippetCreateForm{
		Expires: 365,
	}

	app.render(w, http.StatusOK, "create.html", data)
}

func (app *Application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	// 1. parse + get form body data
	var form snippetCreateForm

	err := app.decodePostForm(r, &form)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// 2. validate form data
	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "create.html", data)
		return
	}

	// 3. insert form to database + response
	id, err := app.Snippets.Insert(form.Title, form.Content, form.Expires)

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
