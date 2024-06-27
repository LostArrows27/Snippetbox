package rest

import "net/http"

type HandlerFunc http.HandlerFunc

type RestAPI struct {
	MUX *http.ServeMux
}

func (rest *RestAPI) Get(path string, handler HandlerFunc, config ...string) {
	handlerGet := handlerMethod(http.MethodGet, handler, path, config...)

	rest.MUX.HandleFunc(path, handlerGet)
}

func (rest *RestAPI) Post(path string, handler HandlerFunc, config ...string) {
	postHandler := handlerMethod(http.MethodPost, handler, path, config...)

	rest.MUX.HandleFunc(path, postHandler)
}
