package handler

import (
	"net/http"

	"github.com/LostArrows27/snippetbox/pkg/rest"
)

func (app *Application) RoutesHandler(restMux rest.RestAPI) *http.ServeMux {
	restMux.Get("/static/", app.StaticFileHanlder)
	restMux.Get("/", app.HomeHandler, "fixed")
	restMux.Get("/snippet/view", app.ViewSnippetHandler)
	restMux.Post("/snippet/create", app.CreateSnippetHanlder)
	return restMux.MUX
}
