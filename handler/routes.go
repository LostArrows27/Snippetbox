package handler

import (
	"net/http"

	"github.com/LostArrows27/snippetbox/pkg/rest"
)

// secureheaders -> middleware -> servermux -> handler
func (app *Application) RoutesHandler(restMux rest.RestAPI) http.Handler {
	restMux.Get("/static/", app.StaticFileHanlder)
	restMux.Get("/", app.HomeHandler, "fixed")
	restMux.Get("/snippet/view", app.ViewSnippetHandler)
	restMux.Post("/snippet/create", app.CreateSnippetHanlder)
	return secureHeaders(restMux.MUX)
}
