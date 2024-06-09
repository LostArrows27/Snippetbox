package htmlParse

import (
	"net/http"
	"text/template"

	"github.com/LostArrows27/snippetbox/pkg/logger"
)

func parseHTML(w http.ResponseWriter, path []string) *template.Template {
	ts, err := template.ParseFiles(path...)

	if err != nil {
		logger.Error(err)
		http.Error(w, "Interal Server Error", 500)
		return nil
	}

	return ts

}

func ExecuteHTML(w http.ResponseWriter, base string, path []string) {
	ts := parseHTML(w, path)

	err := ts.ExecuteTemplate(w, base, nil)
	if err != nil {
		logger.Error(err)
		http.Error(w, "Internal Server Error", 500)
	}
}
