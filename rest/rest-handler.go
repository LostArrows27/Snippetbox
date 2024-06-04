package rest

import "net/http"

func HandlerMethod(method string, handler HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(403)
			w.Write([]byte("Method Not Allowed"))
			return
		}

		handler(w, r)
	}
}
