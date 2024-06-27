package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

// secureheaders -> middleware -> servermux -> handler
func (app *Application) RoutesHandler() http.Handler {

	// 1. configure router
	router := httprouter.New()

	// 2. not found router
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.render(w, http.StatusOK, "not-found.html", &templateData{
			Title: "404 - Not Found",
		})
	})

	// 3. main router + handler
	router.HandlerFunc(http.MethodGet, "/", app.snippetHomeView)
	router.HandlerFunc(http.MethodGet, "/static/*filepath", app.fileHandler)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreateForm)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	// 4. chain middleware
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
