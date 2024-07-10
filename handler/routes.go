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

	// 3. config session manager middleware

	sessionStore := app.SessionManager

	// 4. main router + handler
	router.HandlerFunc(http.MethodGet, "/static/*filepath", app.fileHandler)
	router.Handler(http.MethodGet, "/", sessionStore.LoadAndSave(http.HandlerFunc(app.snippetHomeView)))
	router.Handler(http.MethodGet, "/snippet/view/:id", sessionStore.LoadAndSave(http.HandlerFunc(app.snippetView)))
	router.Handler(http.MethodGet, "/snippet/create", sessionStore.LoadAndSave(http.HandlerFunc(app.snippetCreateForm)))
	router.Handler(http.MethodPost, "/snippet/create", sessionStore.LoadAndSave(http.HandlerFunc(app.snippetCreatePost)))

	// 5. chain middleware
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
