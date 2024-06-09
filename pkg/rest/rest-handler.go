package rest

import "net/http"

// fixed is for "fixed URL path"
// example: / -> won't match with /**
// example: /view/ -> won't match with /view/**

func handlerMethod(method string, handler HandlerFunc, path string, config ...string) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.Header().Set("Allow", method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
