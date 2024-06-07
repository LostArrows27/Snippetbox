package htmlParse

import (
	"log"
	"net/http"
	"text/template"
)

func parseHTML(w http.ResponseWriter, path string) *template.Template {
	ts, err := template.ParseFiles(path)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Interal Server Error", 500)
		return nil
	}

	return ts

}

func ExecuteHTML(w http.ResponseWriter, path string) {
	ts := parseHTML(w, path)

	err := ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}
