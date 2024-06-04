package rest

import "net/http"

// fixed is for "fixed URL path"
// example: / -> won't match with /**
// example: /view/ -> won't match with /view/**

func HandlerMethod(method string, handler HandlerFunc, path string, config ...string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(403)
			w.Write([]byte("Method Not Allowed"))
			return
		}

		isFixed := false

		for _, item := range config {
			if item == "fixed" {
				isFixed = true
				break
			}
		}

		if isFixed && r.URL.Path != path {
			http.NotFound(w, r)
			return
		}

		handler(w, r)
	}
}
