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

	// 3. config unprotected route + session manager middleware

	dynamic := alice.New(app.SessionManager.LoadAndSave, noSurf)

	router.HandlerFunc(http.MethodGet, "/static/*filepath", app.fileHandler)
	router.Handler(http.MethodGet, "/", dynamic.Then(http.HandlerFunc(app.snippetHomeView)))
	router.Handler(http.MethodGet, "/snippet/view/:id", dynamic.Then(http.HandlerFunc(app.snippetView)))
	router.Handler(http.MethodGet, "/user/signup", dynamic.Then(http.HandlerFunc(app.userSignup)))
	router.Handler(http.MethodPost, "/user/signup", dynamic.Then(http.HandlerFunc(app.userSignupPost)))
	router.Handler(http.MethodGet, "/user/login", dynamic.Then(http.HandlerFunc(app.userLogin)))
	router.Handler(http.MethodPost, "/user/login", dynamic.Then(http.HandlerFunc(app.userLoginPost)))

	// 4. config protected route + session manager middleware

	protected := dynamic.Append(app.requiredAuthentication)

	router.Handler(http.MethodGet, "/snippet/create", protected.Then(http.HandlerFunc(app.snippetCreateForm)))
	router.Handler(http.MethodPost, "/snippet/create", protected.Then(http.HandlerFunc(app.snippetCreatePost)))
	router.Handler(http.MethodPost, "/user/logout", protected.Then(http.HandlerFunc(app.userLogoutPost)))

	// 5. chain middleware
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
