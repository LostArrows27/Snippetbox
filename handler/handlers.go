package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	cnst "github.com/LostArrows27/snippetbox/internal/const"
	ipaddress "github.com/LostArrows27/snippetbox/pkg/ip-address"
	"github.com/LostArrows27/snippetbox/pkg/logger"
)

type Application struct {
	ErrorLog logger.CustomLogger
	InfoLog  logger.CustomLogger
}

func (app *Application) HomeHandler(w http.ResponseWriter, r *http.Request) {
	ipaddress.LogRequestIP("/", r)

	ts, err := template.ParseFiles(cnst.HomeHTMLLists...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, cnst.HomeBase, cnst.HomeHTMLLists)

	if err != nil {
		app.serverError(w, err)
	}
}

func (app *Application) CreateSnippetHanlder(w http.ResponseWriter, r *http.Request) {
	ipaddress.LogRequestIP("/snippet/create", r)
	w.Write([]byte("Create snippet"))
}

func (app *Application) ViewSnippetHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	ipaddress.LogRequestIP("/snippet/view?id="+idStr, r)

	if err != nil || id < 0 {
		app.notFound(w)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *Application) StaticFileHanlder(w http.ResponseWriter, r *http.Request) {
	fileServer := http.FileServer(http.Dir(cnst.StaticFileDir))

	handlerFile := http.StripPrefix("/static", fileServer)

	handlerFile.ServeHTTP(w, r)
}
