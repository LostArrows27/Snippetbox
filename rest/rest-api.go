package rest

import "net/http"

type HandlerFunc http.HandlerFunc

type RestAPI struct {
	MUX *http.ServeMux
}

func (rest *RestAPI) Get(path string, handler HandlerFunc) {
	handlerGet := HandlerMethod("GET", handler)

	rest.MUX.HandleFunc(path, handlerGet)
}

func (rest *RestAPI) Post(path string, handler HandlerFunc) {
	postHandler := HandlerMethod("POST", handler)

	rest.MUX.HandleFunc(path, postHandler)

}
